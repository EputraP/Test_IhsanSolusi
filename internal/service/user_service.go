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
)

type UserService interface {
	CreateUser(input *dto.CreateUserBody) (*dto.CreateUserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

type UserServiceConfig struct {
	UserRepo repository.UserRepository
}

func NewUserService(config UserServiceConfig) UserService {
	return &userService{
		userRepo: config.UserRepo,
	}
}
func (s *userService) CreateUser(input *dto.CreateUserBody) (*dto.CreateUserResponse, error) {

	checker, err := s.userRepo.CheckUserByNoHPOrNIK(&model.User{NIK: input.NIK, NoHP: input.NoHP})
	if err != nil {
		return nil, errs.ErrorCheckingUser
	}

	if checker != nil {
		if checker.NIK == input.NIK {
			return nil, errs.ErrorNIKAlreadyUsed
		}
		if checker.NoHP == input.NoHP {
			return nil, errs.ErrorNoHPAlreadyUsed
		}
	}

	noRekening := generate12DigitNumber()

	user, err := s.userRepo.CreateUser(&model.User{NoRekening: noRekening, Nama: strings.ToLower(input.Nama), NIK: input.NIK, NoHP: input.NoHP})
	if err != nil {
		return nil, errs.ErrorCreatingUser
	}

	return &dto.CreateUserResponse{
		NoRekening: user.NoRekening,
	}, err
}

func generate12DigitNumber() string {
	max := new(big.Int).SetUint64(999_999_999_999) // Largest 12-digit number
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%012d", n)
}
