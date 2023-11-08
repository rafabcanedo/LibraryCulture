package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
	Idade int64
	Cargo string
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Idade     int64     `json:"idade"`
	Cargo     string    `json:"cargo"`
}
