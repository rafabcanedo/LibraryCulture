package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Author      string
	Description string
	Reserved    bool
	Price       int64
}

type BookResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Title       string    `json:"role"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Reserved    bool      `json:"reserved"`
	Price       int64     `json:"price"`
}
