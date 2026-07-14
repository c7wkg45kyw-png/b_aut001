package database

import (
	"baut001/backend/internal/config"
	"baut001/backend/internal/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Merchant{}, &entity.User{}, &entity.UserScope{})
}

func SeedAdmin(db *gorm.DB, cfg config.Config) error {
	var merchant entity.Merchant
	if err := db.Where("merchant_code = ?", cfg.SeedMerchantID).First(&merchant).Error; err != nil {
		merchant = entity.Merchant{MerchantCode: cfg.SeedMerchantID, MerchantName: cfg.SeedMerchantName, Status: "active"}
		if err := db.Create(&merchant).Error; err != nil {
			return err
		}
	}
	var user entity.User
	if err := db.Where("username = ?", cfg.SeedAdminUsername).First(&user).Error; err == nil {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(cfg.SeedAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user = entity.User{MerchantID: cfg.SeedMerchantID, Username: cfg.SeedAdminUsername, PasswordHash: string(hash), DisplayName: cfg.SeedAdminName, UserType: "USER", ClientID: cfg.SeedClientID, Status: "active"}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	for _, scope := range cfg.SeedScopes {
		if err := db.Create(&entity.UserScope{UserID: user.ID.String(), Scope: scope}).Error; err != nil {
			return err
		}
	}
	return nil
}
