package dto

import "models"

type WalletRes struct {
	UserID   int    `json:"user_id"`
	WalletID int    `json:"wallet_id"`
	Email    string `json:"email"`
	Balance  int    `json:"balance"`
}

func (wr *WalletRes) FromWallet(w *models.Wallet) *WalletRes {
	return &WalletRes{
		UserID:   w.UserID,
		Balance:  w.Balance,
		WalletID: int(w.ID),
		Email:    w.User.Email,
	}
}
