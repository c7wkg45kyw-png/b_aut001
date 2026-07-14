package middleware

import (
	"net/http"
	"strings"

	"baut001/backend/internal/model"
	"baut001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

const AuthContextKey = "auth_context"

func Auth(authUsecase *usecase.AuthUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			abort(c, http.StatusUnauthorized, "UNAUTHORIZED", "missing bearer token")
			return
		}
		ctx, err := authUsecase.ParseToken(strings.TrimSpace(strings.TrimPrefix(header, "Bearer ")))
		if err != nil {
			abort(c, http.StatusUnauthorized, "UNAUTHORIZED", err.Error())
			return
		}
		if ctx.MerchantID == "" {
			abort(c, http.StatusUnauthorized, "UNAUTHORIZED", "merchant_id claim is required")
			return
		}
		c.Set(AuthContextKey, ctx)
		c.Next()
	}
}

func RequireScope(scope string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := CurrentAuth(c)
		for _, current := range ctx.Scopes {
			if current == scope || current == "auth:*" {
				c.Next()
				return
			}
		}
		abort(c, http.StatusForbidden, "FORBIDDEN", "missing scope: "+scope)
	}
}

func CurrentAuth(c *gin.Context) model.AuthContext {
	value, _ := c.Get(AuthContextKey)
	ctx, _ := value.(model.AuthContext)
	return ctx
}

func abort(c *gin.Context, status int, code, message string) {
	c.AbortWithStatusJSON(status, model.ErrorResponse{Success: false, Code: code, Message: message})
}
