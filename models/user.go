package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"-"`
}
