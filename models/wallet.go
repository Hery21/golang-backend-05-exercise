package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model `json:"-"`
	UserID     int `json:"user_id"`
	Balance    int `json:"Balance"`
	User       User
}
