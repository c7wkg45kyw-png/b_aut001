package main

import (
	"baut001/backend/internal/config"
	"baut001/backend/internal/database"
	"baut001/backend/internal/handler"
	"baut001/backend/internal/logger"
	"baut001/backend/internal/repository"
	"baut001/backend/internal/route"
	"baut001/backend/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.AppEnv)
	defer log.Sync()

	db, err := database.Connect(cfg.DatabaseDSN)
	if err != nil {
		log.Fatal("connect database", zap.Error(err))
	}
	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("migrate database", zap.Error(err))
	}
	if err := database.SeedAdmin(db, cfg); err != nil {
		log.Fatal("seed admin", zap.Error(err))
	}

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo, cfg)
	route.Register(r, cfg, authUsecase, route.Handlers{Auth: handler.NewAuthHandler(authUsecase)})

	log.Info("BAUT001 backend started", zap.String("port", cfg.AppPort))
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("server stopped", zap.Error(err))
	}
}
