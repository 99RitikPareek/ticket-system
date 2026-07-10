package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`

	Status string `json:"status"`

	UserID uint `json:"user_id"`
}