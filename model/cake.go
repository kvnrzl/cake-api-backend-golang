package model

import "time"

type Cake struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Rating      float32   `json:"rating" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
