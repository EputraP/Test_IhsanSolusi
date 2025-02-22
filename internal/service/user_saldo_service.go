package service

import (
	"strconv"

	"github.com/EputraP/Test_IhsanSolusi/internal/dto"
	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
	"github.com/EputraP/Test_IhsanSolusi/internal/model"
	"github.com/EputraP/Test_IhsanSolusi/internal/repository"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/logger"
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

	logger.Info("Starting GetUserSaldoByNoRek", "input", *noRek)

	checkNoRek, err := s.userRepo.CheckUserByNoRek(&model.User{NoRekening: *noRek})
	if err != nil || checkNoRek == nil {
		logger.Error("No Rekening is invalid or does not exist", "error", err)
		return nil, errs.InvalidNoRek
	}

	userSaldo, err := s.userSaldoRepo.GetUserSaldoById(&model.UserSaldo{NoRekening: *noRek})
	if err != nil {
		logger.Error("Error getting user saldo", "error", err)
		return nil, errs.ErrorGettingUserSaldo
	}

	logger.Debug("User saldo retrieved successfully", "noRekening", *noRek, "saldo", userSaldo.Saldo)

	return &dto.CurrentBalanceResponse{
		SaldoSaatIni: userSaldo.Saldo,
	}, err
}

func (s *userSaldoService) TabungTarikSaldo(mode string, input *dto.TransactionBody) (*dto.CurrentBalanceResponse, error) {

	logger.Info("Starting TabungTarikSaldo", "mode", mode, "input", *input)

	checkNoRek, err := s.userRepo.CheckUserByNoRek(&model.User{NoRekening: input.NoRekening})
	if err != nil || checkNoRek == nil {
		logger.Error("No Rekening is invalid or does not exist", "error", err)
		return nil, errs.InvalidNoRek
	}

	userSaldo, err := s.userSaldoRepo.GetUserSaldoById(&model.UserSaldo{NoRekening: input.NoRekening})
	if err != nil {
		logger.Error("Error getting user saldo", "error", err)
		return nil, errs.ErrorGettingUserSaldo
	}

	currentSaldo, err := strconv.Atoi(userSaldo.Saldo)
	if err != nil {
		logger.Error("Error converting saldo to integer", "error", err)
		return nil, errs.ErrorStringIntConversion
	}

	topUpSaldo, err := strconv.Atoi(input.Nominal)
	if err != nil {
		logger.Error("Error converting nominal to integer", "error", err)
		return nil, errs.ErrorStringIntConversion
	}

	if mode == "tarik" {
		topUpSaldo = -topUpSaldo
	}

	newSaldo := topUpSaldo + currentSaldo
	if newSaldo < 0 {
		logger.Error("Insufficient balance for the transaction", "noRekening", input.NoRekening, "currentSaldo", currentSaldo, "topUpSaldo", topUpSaldo)
		return nil, errs.InsufficientBalance
	}

	updateSaldo, err := s.userSaldoRepo.UpdateUserSaldo(&model.UserSaldo{NoRekening: input.NoRekening, Saldo: strconv.Itoa(newSaldo)})
	if err != nil {
		logger.Error("Error updating user saldo", "error", err)
		return nil, err
	}

	logger.Debug("User saldo updated successfully", "noRekening", input.NoRekening, "newSaldo", updateSaldo.Saldo)

	return &dto.CurrentBalanceResponse{
		SaldoSaatIni: updateSaldo.Saldo,
	}, err
}
