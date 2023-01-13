package dto

type TransferReq struct {
	Amount      int    `json:"amount"`
	WalletID    int    `json:"wallet_id"`
	TargetID    int    `json:"target_wallet_id"`
	Description string `json:"description"`
}
