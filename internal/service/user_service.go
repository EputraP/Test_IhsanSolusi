package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/EputraP/Test_IhsanSolusi/internal/dto"
	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
	"github.com/EputraP/Test_IhsanSolusi/internal/model"
	"github.com/EputraP/Test_IhsanSolusi/internal/repository"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/logger"
)

type UserService interface {
	CreateUser(input *dto.CreateUserBody) (*dto.CreateUserResponse, error)
}

type userService struct {
	userRepo      repository.UserRepository
	userSaldoRepo repository.UserSaldoRepository
}

type UserServiceConfig struct {
	UserRepo      repository.UserRepository
	UserSaldoRepo repository.UserSaldoRepository
}

func NewUserService(config UserServiceConfig) UserService {
	return &userService{
		userRepo:      config.UserRepo,
		userSaldoRepo: config.UserSaldoRepo,
	}
}

func (s *userService) CreateUser(input *dto.CreateUserBody) (*dto.CreateUserResponse, error) {

	logger.Info("Starting CreateUser", "input", *input)

	checker, err := s.userRepo.CheckUserByNoHPOrNIK(&model.User{NIK: input.NIK, NoHP: input.NoHP})
	if err != nil {
		logger.Error("Error checking user by NIK or NoHP", "error", err)
		return nil, errs.ErrorCheckingUser
	}

	if checker != nil {
		if checker.NIK == input.NIK {
			logger.Error("NIK already used", "NIK", input.NIK)
			return nil, errs.ErrorNIKAlreadyUsed
		}
		if checker.NoHP == input.NoHP {
			logger.Error("NoHP already used", "NoHP", input.NoHP)
			return nil, errs.ErrorNoHPAlreadyUsed
		}
	}

	noRekening := generate12DigitNumber()
	logger.Info("Generated new NoRekening", "NoRekening", noRekening)

	user, err := s.userRepo.CreateUser(&model.User{NoRekening: noRekening, Nama: strings.ToLower(input.Nama), NIK: input.NIK, NoHP: input.NoHP})
	if err != nil {
		logger.Error("Error creating user", "error", err)
		return nil, errs.ErrorCreatingUser
	}

	logger.Info("User created successfully", "NoRekening", user.NoRekening)

	_, err = s.userSaldoRepo.CreateUserSaldo(&model.UserSaldo{NoRekening: user.NoRekening, Saldo: "0"})
	if err != nil {
		logger.Error("Error creating user saldo", "error", err)
		return nil, errs.ErrorCreatingUserSaldo
	}

	logger.Info("User saldo created successfully", "NoRekening", user.NoRekening)

	return &dto.CreateUserResponse{
		NoRekening: user.NoRekening,
	}, nil
}

func generate12DigitNumber() string {
	max := new(big.Int).SetUint64(999_999_999_999) // Largest 12-digit number
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		logger.Error("Error generating 12 digit number", "error", err)
		panic(err)
	}
	return fmt.Sprintf("%012d", n)
}
