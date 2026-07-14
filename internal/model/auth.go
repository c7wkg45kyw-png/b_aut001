package model

import "time"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type TokenPayload struct {
	Issuer     string   `json:"iss"`
	Subject    string   `json:"sub"`
	Audience   string   `json:"aud"`
	ExpiresAt  int64    `json:"exp"`
	IssuedAt   int64    `json:"iat"`
	MerchantID string   `json:"merchant_id"`
	ClientID   string   `json:"client_id"`
	Type       string   `json:"type"`
	Scopes     []string `json:"scopes"`
}

type UserProfile struct {
	ID          string    `json:"id"`
	MerchantID  string    `json:"merchant_id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	ClientID    string    `json:"client_id"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Scopes      []string  `json:"scopes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AuthContext struct {
	UserID     string
	MerchantID string
	Scopes     []string
}
