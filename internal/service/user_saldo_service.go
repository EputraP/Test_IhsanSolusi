package service

import (
	"github.com/EputraP/Test_IhsanSolusi/internal/dto"
	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
	"github.com/EputraP/Test_IhsanSolusi/internal/model"
	"github.com/EputraP/Test_IhsanSolusi/internal/repository"
)

type UserSaldoService interface {
}

type userSaldoService struct {
	userSaldoRepo repository.UserSaldoRepository
	userRepo      repository.UserRepository
}

type UserSaldoServiceConfig struct {
	UserSaldoRepo repository.UserSaldoRepository
	UserRepo      repository.UserRepository
}

func NewUserSaldoService(config UserSaldoServiceConfig) UserSaldoService {
	return &userSaldoService{
		userSaldoRepo: config.UserSaldoRepo,
		userRepo:      config.UserRepo,
	}
}

func (s *userSaldoService) GetUserSaldoByNoRek(noRek *string) (*dto.CurrentBalanceResponse, error) {

	checkNoRek, err := s.userRepo.CheckUserByNoRek(&model.User{NoRekening: *noRek})
	if err != nil || checkNoRek == nil {
		return nil, errs.InvalidNoRek
	}

	userSaldo, err := s.userSaldoRepo.GetUserSaldoById(&model.UserSaldo{NoRekening: *noRek})
	if err != nil {
		return nil, errs.ErrorGettingUserSaldo
	}

	return &dto.CurrentBalanceResponse{
		SaldoSaatIni: userSaldo.Saldo,
	}, err
}
