package handler

import (
	"net/http"

	"baut001/backend/internal/middleware"
	"baut001/backend/internal/model"
	"baut001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "VALIDATION_ERROR", err.Error())
		return
	}
	result, err := h.usecase.Login(req)
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "login success", result)
}

func (h *AuthHandler) Me(c *gin.Context) {
	result, err := h.usecase.Me(middleware.CurrentAuth(c))
	if err != nil {
		handleError(c, err)
		return
	}
	ok(c, "current user", result)
}
