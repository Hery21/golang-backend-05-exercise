package repositories

import (
	"errors"
	" hery-ciaputra/assignment-05-golang-backend/httperror"
	" hery-ciaputra/assignment-05-golang-backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	MatchingCredential(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) *userRepository {
	return &userRepository{db: c.DB}
}

func (w *userRepository) MatchingCredential(email string) (*models.User, error) {
	var user *models.User

	res := w.db.Where("email = ?", email).First(&user)

	isNotFound := errors.Is(res.Error, gorm.ErrRecordNotFound)
	if isNotFound {
		return nil, httperror.BadRequestError("Invalid email or password", "INVALID_LOGIN")
	}
	return user, nil
}
