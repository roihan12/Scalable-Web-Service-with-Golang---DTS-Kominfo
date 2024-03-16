package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string
	Description string
	UserID      uint
	User        *User


	
}
