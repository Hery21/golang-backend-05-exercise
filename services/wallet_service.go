package services

import (
	"fmt"
	" hery-ciaputra/assignment-05-golang-backend/dto"
	" hery-ciaputra/assignment-05-golang-backend/models"
	" hery-ciaputra/assignment-05-golang-backend/repositories"
	"time"
)

type WalletService interface {
	GetWallet(id int) (*dto.WalletRes, error)
	TopUpWallet(id int, req *dto.TopUpReq) (*dto.TopUpRes, error)
	TransferWallet(id int, req *dto.TransferReq) (*dto.TransferRes, error)
	GetTransactions(id int, transactionFilter *models.Filter) ([]*dto.TransactionsRes, error)
}

type walletService struct {
	walletRepository repositories.WalletRepository
}

type WSConfig struct {
	WalletRepository repositories.WalletRepository
}

func NewWalletService(w *WSConfig) WalletService {
	return &walletService{
		walletRepository: w.WalletRepository,
	}
}

func (w *walletService) GetWallet(id int) (*dto.WalletRes, error) {
	wallet, err := w.walletRepository.GetWallet(id)
	if err != nil {
		return new(dto.WalletRes), err
	}

	return new(dto.WalletRes).FromWallet(wallet), nil
}

func (w *walletService) TopUpWallet(id int, req *dto.TopUpReq) (*dto.TopUpRes, error) {
	topUpInfo := &models.Transaction{
		Amount:          req.Amount,
		TransactionType: "CREDIT",
		WalletID:        req.WalletID,
		TargetWalletID:  req.TargetID,
		CreatedAt:       time.Now(),
	}

	if req.TargetID == 1001 {
		topUpInfo.Description = "Top Up from Bank Transfer"
	}

	if req.TargetID == 1002 {
		topUpInfo.Description = "Top Up from Cash"
	}

	if req.TargetID == 1003 {
		topUpInfo.Description = "Top Up from Credit Card"
	}

	topUp, err := w.walletRepository.TopUpWallet(id, topUpInfo)
	if err != nil {
		return new(dto.TopUpRes), err
	}

	return new(dto.TopUpRes).FromTopUp(topUp), nil
}

func (w *walletService) TransferWallet(id int, req *dto.TransferReq) (*dto.TransferRes, error) {
	transferInfo := &models.Transaction{
		Amount:          req.Amount,
		TransactionType: "DEBIT",
		WalletID:        req.WalletID,
		TargetWalletID:  req.TargetID,
		Description:     req.Description,
		CreatedAt:       time.Now(),
	}

	transfer, err := w.walletRepository.TransferWallet(id, transferInfo)
	if err != nil {
		return new(dto.TransferRes), err
	}

	return new(dto.TransferRes).FromTransfer(transfer), nil
}

func (w *walletService) GetTransactions(id int, transactionFilter *models.Filter) ([]*dto.TransactionsRes, error) {
	filteredTransaction, err := w.walletRepository.GetTransactions(id, transactionFilter)
	var filteredTransactionList []*dto.TransactionsRes

	for i := range filteredTransaction {
		filteredTransactionList = append(filteredTransactionList, new(dto.TransactionsRes).FromTransactions(filteredTransaction[i]))
	}
	fmt.Println(filteredTransactionList)

	return filteredTransactionList, err
}
