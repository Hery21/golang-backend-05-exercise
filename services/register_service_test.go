package services_test

import (
	" hery-ciaputra/assignment-05-golang-backend/dto"
	" hery-ciaputra/assignment-05-golang-backend/models"
	"gorm.io/gorm"
	"testing"
)

var (
	registerReq = dto.RegisterReq{
		Email:    "user1@mail.com",
		Name:     "user1",
		Password: "",
	}

	user = models.User{
		Model:    gorm.Model{},
		ID:       1,
		Name:     "user1",
		Email:    "user1@mail.com",
		Password: "",
	}

	registerRes = dto.RegisterRes{
		ID:    1,
		Name:  "user1",
		Email: "user1@mail.com",
	}
)

func Test_registerService_Register(t *testing.T) {
	//t.Run("Should return status 200 and registered user info", func(t *testing.T) {
	//	mockRepository := new(mocks.RegisterRepository)
	//	config := server.RouterConfig{mockRepository}
	//	s := server.NewLoanRequest(&config)
	//	loanRequestInput := loanRequest
	//	loanRequestInput.ID = 0
	//	mockRepository.On("Create", &loanRequestInput).Return(&loanRequest, nil)
	//	payload := dto.PostLoanReq{
	//		UserID: 1,
	//		BookID: 1,
	//	}
	//	expectedLoanRequestRes := &loanRequestRes
	//
	//	actualLoanRequestRes, err := s.Create(&payload)
	//
	//	assert.Nil(t, err)
	//	assert.Equal(t, expectedLoanRequestRes, actualLoanRequestRes)
	//})
}
