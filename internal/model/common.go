package model

import "time"

type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AuditFields struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
