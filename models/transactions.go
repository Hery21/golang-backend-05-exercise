package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model      `json:"-"`
	Amount          int       `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	WalletID        int       `json:"wallet_id"`
	TargetWalletID  int       `json:"target_wallet_id"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
}
