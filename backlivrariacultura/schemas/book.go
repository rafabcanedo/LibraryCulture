package schemas

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	author      string
	description string
	reserved    bool
	price       int64
}
