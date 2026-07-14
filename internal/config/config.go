package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppEnv             string
	AppPort            string
	DatabaseDSN        string
	JWTSecret          string
	JWTIssuer          string
	JWTDefaultAudience string
	JWTTTLMinutes      int
	CORSAllowOrigins   []string
	SeedAdminUsername  string
	SeedAdminPassword  string
	SeedAdminName      string
	SeedMerchantID     string
	SeedMerchantName   string
	SeedClientID       string
	SeedScopes         []string
}

func Load() Config {
	loadDotEnv(".env")
	return Config{
		AppEnv:             getenv("APP_ENV", "development"),
		AppPort:            getenv("APP_PORT", "8081"),
		DatabaseDSN:        getenv("DATABASE_DSN", "host=localhost user=postgres password=postgres dbname=auth_001 port=5432 sslmode=disable TimeZone=Asia/Bangkok"),
		JWTSecret:          getenv("JWT_SECRET", "change-me-please-use-32-plus-chars"),
		JWTIssuer:          getenv("JWT_ISSUER", "https://global-commerce.com"),
		JWTDefaultAudience: getenv("JWT_DEFAULT_AUDIENCE", "sku_module"),
		JWTTTLMinutes:      getenvInt("JWT_TTL_MINUTES", 1440),
		CORSAllowOrigins:   splitCSV(getenv("CORS_ALLOW_ORIGINS", "*")),
		SeedAdminUsername:  getenv("SEED_ADMIN_USERNAME", "admin"),
		SeedAdminPassword:  getenv("SEED_ADMIN_PASSWORD", "admin1234"),
		SeedAdminName:      getenv("SEED_ADMIN_DISPLAY_NAME", "System Admin"),
		SeedMerchantID:     getenv("SEED_MERCHANT_ID", "mch_555666777"),
		SeedMerchantName:   getenv("SEED_MERCHANT_NAME", "Demo Merchant"),
		SeedClientID:       getenv("SEED_CLIENT_ID", "vendor_portal_web"),
		SeedScopes:         splitCSV(getenv("SEED_SCOPES", "sku:read,sku:update")),
	}
}

func getenv(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func getenvInt(key string, fallback int) int {
	value, err := strconv.Atoi(getenv(key, ""))
	if err != nil || value <= 0 {
		return fallback
	}
	return value
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func loadDotEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}
		pair := strings.SplitN(line, "=", 2)
		key := strings.TrimSpace(pair[0])
		value := strings.Trim(strings.TrimSpace(pair[1]), "\"")
		if os.Getenv(key) == "" {
			_ = os.Setenv(key, value)
		}
	}
}
