package route

import (
	"net/http"

	"baut001/backend/internal/config"
	"baut001/backend/internal/handler"
	"baut001/backend/internal/middleware"
	"baut001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth *handler.AuthHandler
}

func Register(router *gin.Engine, cfg config.Config, authUsecase *usecase.AuthUsecase, handlers Handlers) {
	router.Use(middleware.CORS(cfg))
	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"success": true, "message": "healthy"}) })
	router.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")
	router.GET("/docs/swagger.html", swaggerHTML)

	api := router.Group("/api/v1")
	api.POST("/auth/login", handlers.Auth.Login)

	protected := api.Group("")
	protected.Use(middleware.Auth(authUsecase))
	protected.GET("/auth/me", handlers.Auth.Me)
}

func swaggerHTML(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, `<!doctype html><html><head><title>BAUT001 Swagger</title><link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css"></head><body><div id="swagger-ui"></div><script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script><script>window.onload=()=>SwaggerUIBundle({url:'/docs/openapi.yaml',dom_id:'#swagger-ui'});</script></body></html>`)
}
