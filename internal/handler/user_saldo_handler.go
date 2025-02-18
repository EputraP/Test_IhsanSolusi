package handler

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/service"
)

type UserSaldoHandler struct {
	userSaldoService service.UserSaldoService
}

type UserSaldoHandlerConfig struct {
	UserSaldoService service.UserSaldoService
}

func NewUserSaldoHandler(config UserSaldoHandlerConfig) *UserSaldoHandler {
	return &UserSaldoHandler{
		userSaldoService: config.UserSaldoService,
	}
}
