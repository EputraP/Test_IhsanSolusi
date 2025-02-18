package service

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/repository"
)

type UserSaldoService interface {
}

type userSaldoService struct {
	userSaldoRepo repository.UserSaldoRepository
}

type UserSaldoServiceConfig struct {
	UserSaldoRepo repository.UserSaldoRepository
}

func NewUserSaldoService(config UserSaldoServiceConfig) UserSaldoService {
	return &userSaldoService{
		userSaldoRepo: config.UserSaldoRepo,
	}
}
