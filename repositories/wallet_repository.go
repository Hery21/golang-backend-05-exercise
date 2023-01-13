package repositories

import (
	"fmt"
	" hery-ciaputra/assignment-05-golang-backend/httperror"
	" hery-ciaputra/assignment-05-golang-backend/models"
	"gorm.io/gorm"
	"strconv"
)

type WalletRepository interface {
	GetWallet(id int) (*models.Wallet, error)
	TopUpWallet(id int, up *models.Transaction) (*models.Transaction, error)
	TransferWallet(id int, tu *models.Transaction) (*models.Transaction, error)
	GetTransactions(id int, filter *models.Filter) ([]*models.Transaction, error)
}

type walletRepository struct {
	db *gorm.DB
}

type WRConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(c *WRConfig) *walletRepository {
	return &walletRepository{db: c.DB}
}

func (w *walletRepository) GetWallet(id int) (*models.Wallet, error) {
	var wallet *models.Wallet

	res := w.db.Joins("User").Where("user_id = ?", id).First(&wallet)
	fmt.Println(res)

	if res.Error != nil {
		return nil, res.Error
	}

	return wallet, nil
}

func (w *walletRepository) TopUpWallet(id int, tu *models.Transaction) (*models.Transaction, error) {
	var wallet *models.Wallet

	errWallet := w.db.Where("user_id = ?", id).First(&wallet).UpdateColumn("balance", gorm.Expr("balance + ?", tu.Amount))
	if errWallet.Error != nil {
		return nil, errWallet.Error
	}

	errTransaction := w.db.Create(&tu)
	if errTransaction.Error != nil {
		return nil, errTransaction.Error
	}

	return tu, nil
}

func (w *walletRepository) TransferWallet(id int, tu *models.Transaction) (*models.Transaction, error) {
	var wallet *models.Wallet
	var targetWallet *models.Wallet

	errWallet := w.db.Where("user_id = ?", id).First(&wallet).UpdateColumn("balance", gorm.Expr("balance + ?", tu.Amount))
	if errWallet.Error != nil {
		return nil, errWallet.Error
	}

	errTarget := w.db.Where("id = ?", tu.TargetWalletID).First(&targetWallet).UpdateColumn("balance", gorm.Expr("balance - ?", tu.Amount))
	if errTarget.Error != nil {
		return nil, errTarget.Error
	}

	errTransaction := w.db.Create(&tu)
	if errTransaction.Error != nil {
		return nil, errTransaction.Error
	}

	return tu, nil
}

func (w *walletRepository) GetTransactions(id int, filter *models.Filter) ([]*models.Transaction, error) {
	var wallet *models.Wallet
	var filteredTransaction []*models.Transaction

	res := w.db.Select("id").Where("user_id = ?", id).First(&wallet)
	if res.Error != nil {
		return nil, httperror.BadRequestError("Wallet not found", "WALLET_NOT_FOUND")
	}

	orderBy := filter.SortBy + " " + filter.Sort

	limit, _ := strconv.Atoi(filter.Limit)

	if filter.Limit == "" {
		limit = 10
	}

	res2 := w.db.Table("transactions").Where("wallet_id = ? and description iLIKE ?", wallet.ID, "%"+filter.Search+"%").Order(orderBy).Limit(limit).Scan(&filteredTransaction)

	if res2.Error != nil {
		return nil, res2.Error
	}

	return filteredTransaction, nil
}
