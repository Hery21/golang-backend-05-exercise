package server

import (
	"fmt"
	" hery-ciaputra/assignment-05-golang-backend/config"
	" hery-ciaputra/assignment-05-golang-backend/db"
	" hery-ciaputra/assignment-05-golang-backend/repositories"
	" hery-ciaputra/assignment-05-golang-backend/services"
)

func Init() {
	walletRepository := repositories.NewWalletRepository(&repositories.WRConfig{DB: db.Get()})
	walletService := services.NewWalletService(&services.WSConfig{WalletRepository: walletRepository})

	registerRepository := repositories.NewRegisterRepository(&repositories.RRConfig{DB: db.Get()})
	registerService := services.NewRegisterService(&services.RSConfig{RegisterRepository: registerRepository})

	userRepository := repositories.NewUserRepository(&repositories.URConfig{DB: db.Get()})
	authService := services.NewAuthService(
		&services.AuthSConfig{
			UserRepository: userRepository,
			AppConfig:      config.Config,
		})

	router := NewRouter(&RouterConfig{
		WalletService:   walletService,
		AuthService:     authService,
		RegisterService: registerService,
	})

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("server Error: ", err)
	}
}
