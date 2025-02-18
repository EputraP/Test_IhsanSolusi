package service

import (
	"strconv"

	"github.com/EputraP/Test_IhsanSolusi/internal/dto"
	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
	"github.com/EputraP/Test_IhsanSolusi/internal/model"
	"github.com/EputraP/Test_IhsanSolusi/internal/repository"
)

type UserSaldoService interface {
	GetUserSaldoByNoRek(noRek *string) (*dto.CurrentBalanceResponse, error)
	TabungTarikSaldo(mode string, input *dto.TransactionBody) (*dto.CurrentBalanceResponse, error)
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

func (s *userSaldoService) TabungTarikSaldo(mode string, input *dto.TransactionBody) (*dto.CurrentBalanceResponse, error) {

	checkNoRek, err := s.userRepo.CheckUserByNoRek(&model.User{NoRekening: input.NoRekening})
	if err != nil || checkNoRek == nil {
		return nil, errs.InvalidNoRek
	}

	userSaldo, err := s.userSaldoRepo.GetUserSaldoById(&model.UserSaldo{NoRekening: input.NoRekening})
	if err != nil {
		return nil, errs.ErrorGettingUserSaldo
	}

	currentSaldo, err := strconv.Atoi(userSaldo.Saldo)
	if err != nil {
		return nil, errs.ErrorStringIntConvertion
	}
	topUpSaldo, err := strconv.Atoi(input.Nominal)
	if err != nil {
		return nil, errs.ErrorStringIntConvertion
	}

	if mode == "tarik" {
		topUpSaldo = -topUpSaldo
	}

	newSaldo := topUpSaldo + currentSaldo
	if newSaldo < 0 {
		return nil, errs.InsufficientBalance
	}

	updateSaldo, err := s.userSaldoRepo.UpdateUserSaldo(&model.UserSaldo{NoRekening: input.NoRekening, Saldo: strconv.Itoa(newSaldo)})

	return &dto.CurrentBalanceResponse{
		SaldoSaatIni: updateSaldo.Saldo,
	}, err
}
