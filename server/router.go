package server

import (
	" hery-ciaputra/assignment-05-golang-backend/dto"
	" hery-ciaputra/assignment-05-golang-backend/handlers"
	" hery-ciaputra/assignment-05-golang-backend/middlewares"
	" hery-ciaputra/assignment-05-golang-backend/services"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	WalletService   services.WalletService
	AuthService     services.AuthService
	RegisterService services.RegisterService
}

func NewRouter(c *RouterConfig) *gin.Engine {

	router := gin.Default()

	h := handlers.New(&handlers.HandlerConfig{
		WalletService:   c.WalletService,
		AuthService:     c.AuthService,
		RegisterService: c.RegisterService,
	})
	router.Static("/docs", "swaggerui")

	router.POST("/sign-in", middlewares.RequestValidator(func() any {
		return &dto.SignInReq{}
	}), h.SignIn, middlewares.ErrorHandler)

	router.POST("/register", middlewares.RequestValidator(func() any {
		return &dto.RegisterReq{}
	}), h.Register, middlewares.ErrorHandler)

	router.GET("/wallets", middlewares.AuthorizeJWT, h.GetWallet, middlewares.ErrorHandler)

	router.POST("/wallets/top-up", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.TopUpReq{}
	}), h.TopUpWallet, middlewares.ErrorHandler)

	router.POST("/wallets/transfer", middlewares.AuthorizeJWT, middlewares.RequestValidator(func() any {
		return &dto.TransferReq{}
	}), h.TransferWallet, middlewares.ErrorHandler)

	router.GET("/wallets/transactions", middlewares.AuthorizeJWT, h.GetTransactions)

	return router
}
