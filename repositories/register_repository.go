package repositories

import (
	" hery-ciaputra/assignment-05-golang-backend/httperror"
	" hery-ciaputra/assignment-05-golang-backend/models"
)

type RegisterRepository interface {
	Register(user *models.User) (*models.Register, error)
	RegisterWallet(wallet *models.Wallet) (*models.Wallet, error)
}

type registerRepository struct {
	db *gorm.DB
}

type RRConfig struct {
	DB *gorm.DB
}

func NewRegisterRepository(c *RRConfig) *registerRepository {
	return &registerRepository{db: c.DB}
}

func (r *registerRepository) Register(user *models.User) (*models.Register, error) {
	res := r.db.Select("Name", "Email", "Password").Create(&user)

	if err := res.Error; err != nil {
		return &models.Register{}, httperror.InternalServerError(err.Error())
	}

	registeredUser := &models.Register{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return registeredUser, res.Error
}

func (r *registerRepository) RegisterWallet(wallet *models.Wallet) (*models.Wallet, error) {
	res := r.db.Select("user_id", "balance").Create(&wallet)

	if err := res.Error; err != nil {
		return &models.Wallet{}, httperror.InternalServerError(err.Error())
	}
	return wallet, nil
}
