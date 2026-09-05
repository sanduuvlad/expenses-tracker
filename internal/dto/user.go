package dto

import "time"

type UserResponse struct {
	ID        int64
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
