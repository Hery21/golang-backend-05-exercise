package dto

import (
	"time"

	"models"
)

type TransactionsRes struct {
	Date            time.Time `json:"created_at"`
	WalletID        int       `json:"wallet_id"`
	TargetID        int       `json:"target_id"`
	TransactionType string    `json:"transaction_type"`
	Description     string    `json:"description"`
	Amount          int       `json:"amount"`
}

func (tr *TransactionsRes) FromTransactions(m *models.Transaction) *TransactionsRes {
	return &TransactionsRes{
		Date:            m.CreatedAt,
		WalletID:        m.WalletID,
		TargetID:        m.TargetWalletID,
		TransactionType: m.TransactionType,
		Description:     m.Description,
		Amount:          m.Amount,
	}
}
