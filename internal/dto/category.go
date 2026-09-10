package dto

import "time"

type CategoryResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
