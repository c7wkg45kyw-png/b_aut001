package usecase

import (
	"errors"
	"time"

	"baut001/backend/internal/config"
	"baut001/backend/internal/entity"
	"baut001/backend/internal/model"
	"baut001/backend/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	MerchantID string   `json:"merchant_id"`
	BranchID   string   `json:"branch_id"`
	ClientID   string   `json:"client_id"`
	Type       string   `json:"type"`
	Scopes     []string `json:"scopes"`
	jwt.RegisteredClaims
}

type AuthUsecase struct {
	repo *repository.AuthRepository
	cfg  config.Config
}

func NewAuthUsecase(repo *repository.AuthRepository, cfg config.Config) *AuthUsecase {
	return &AuthUsecase{repo: repo, cfg: cfg}
}

func (u *AuthUsecase) Login(req model.LoginRequest) (model.TokenResponse, error) {
	user, err := u.repo.FindUserByUsername(req.Username)
	if err != nil {
		return model.TokenResponse{}, errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return model.TokenResponse{}, errors.New("invalid username or password")
	}
	scopes, err := u.repo.ListScopes(user.ID.String())
	if err != nil {
		return model.TokenResponse{}, err
	}
	return u.issueToken(user, scopes, u.cfg.JWTDefaultAudience, user.ClientID)
}

func (u *AuthUsecase) Me(ctx model.AuthContext) (model.UserProfile, error) {
	user, err := u.repo.GetUserByID(ctx.UserID)
	if err != nil {
		return model.UserProfile{}, err
	}
	scopes, err := u.repo.ListScopes(user.ID.String())
	if err != nil {
		return model.UserProfile{}, err
	}
	return mapUser(user, scopes), nil
}

func (u *AuthUsecase) ParseToken(tokenText string) (model.AuthContext, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenText, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(u.cfg.JWTSecret), nil
	}, jwt.WithIssuer(u.cfg.JWTIssuer))
	if err != nil || !token.Valid {
		return model.AuthContext{}, errors.New("invalid token")
	}
	return model.AuthContext{UserID: claims.Subject, MerchantID: claims.MerchantID, Scopes: claims.Scopes}, nil
}

func (u *AuthUsecase) issueToken(user entity.User, scopes []string, audience, clientID string) (model.TokenResponse, error) {
	now := time.Now().UTC()
	expires := now.Add(time.Duration(u.cfg.JWTTTLMinutes) * time.Minute)
	claims := Claims{
		MerchantID: user.MerchantID,
		ClientID:   clientID,
		Type:       user.UserType,
		Scopes:     scopes,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    u.cfg.JWTIssuer,
			Subject:   user.ID.String(),
			Audience:  jwt.ClaimStrings{audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expires),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return model.TokenResponse{}, err
	}
	return model.TokenResponse{AccessToken: signed, TokenType: "Bearer"}, nil
}

func mapUser(user entity.User, scopes []string) model.UserProfile {
	return model.UserProfile{ID: user.ID.String(), MerchantID: user.MerchantID, Username: user.Username, DisplayName: user.DisplayName, Email: user.Email, ClientID: user.ClientID, Type: user.UserType, Status: user.Status, Scopes: scopes, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
}
