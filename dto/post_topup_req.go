package dto

type TopUpReq struct {
	Amount   int `json:"amount"`
	WalletID int `json:"wallet_id"`
	TargetID int `json:"target_wallet_id"`
}
