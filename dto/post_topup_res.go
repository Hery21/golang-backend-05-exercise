package dto

import "models"

type TopUpRes struct {
	ID              int    `json:"id"`
	Amount          int    `json:"amount"`
	TransactionType string `json:"transaction_type"`
	WalletID        int    `json:"wallet_id"`
	TargetID        int    `json:"target_wallet_id"`
	Description     string `json:"description"`
}

func (tu *TopUpRes) FromTopUp(t *models.Transaction) *TopUpRes {
	return &TopUpRes{
		ID:              int(t.ID),
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		WalletID:        t.WalletID,
		TargetID:        t.TargetWalletID,
		Description:     t.Description,
	}
}
