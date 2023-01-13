package handlers

import (
	"services"
)

type HandlerConfig struct {
	WalletService   services.WalletService
	AuthService     services.AuthService
	RegisterService services.RegisterService
}

type Handler struct {
	walletService   services.WalletService
	authService     services.AuthService
	registerService services.RegisterService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		walletService:   c.WalletService,
		authService:     c.AuthService,
		registerService: c.RegisterService,
	}
}
