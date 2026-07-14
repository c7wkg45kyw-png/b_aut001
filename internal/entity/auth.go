package entity

type Merchant struct {
	Base
	MerchantCode string `gorm:"size:80;not null;uniqueIndex" json:"merchant_code"`
	MerchantName string `gorm:"size:255;not null" json:"merchant_name"`
	Status       string `gorm:"size:40;not null;default:'active'" json:"status"`
}

type User struct {
	Base
	MerchantID   string `gorm:"size:80;not null;index" json:"merchant_id"`
	Username     string `gorm:"size:120;not null;uniqueIndex" json:"username"`
	PasswordHash string `gorm:"size:255;not null" json:"-"`
	DisplayName  string `gorm:"size:255;not null" json:"display_name"`
	Email        string `gorm:"size:255" json:"email"`
	UserType     string `gorm:"size:40;not null;default:'USER'" json:"type"`
	ClientID     string `gorm:"size:120;not null;default:'vendor_portal_web'" json:"client_id"`
	Status       string `gorm:"size:40;not null;default:'active'" json:"status"`
}

type UserScope struct {
	Base
	UserID string `gorm:"type:uuid;not null;index" json:"user_id"`
	Scope  string `gorm:"size:120;not null;index" json:"scope"`
}
