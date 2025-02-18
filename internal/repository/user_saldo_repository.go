package repository

import (
	"time"

	"github.com/EputraP/Test_IhsanSolusi/internal/model"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/logger"
	"gorm.io/gorm"
)

type UserSaldoRepository interface {
	CreateUserSaldo(inputModel *model.UserSaldo) (*model.UserSaldo, error)
	GetUserSaldoById(inputModel *model.UserSaldo) (*model.UserSaldo, error)
	UpdateUserSaldo(inputModel *model.UserSaldo) (*model.UserSaldo, error)
}

type userSaldoRepository struct {
	db *gorm.DB
}

func NewUserSaldoRepository(db *gorm.DB) UserSaldoRepository {
	return &userSaldoRepository{
		db: db,
	}
}

func (r *userSaldoRepository) CreateUserSaldo(inputModel *model.UserSaldo) (*model.UserSaldo, error) {

	sqlScript := `INSERT INTO user_saldo (no_rekening , saldo , created_at) 
				VALUES (?,?,?) 
				RETURNING id;`

	logger.Info("Executing CreateUserSaldo SQL query", "input", *inputModel)

	res := r.db.Raw(sqlScript, inputModel.NoRekening, inputModel.Saldo, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("Failed to create user saldo in CreateUserSaldo", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("Successfully created user saldo", "output", *inputModel)

	return inputModel, nil
}

func (r *userSaldoRepository) GetUserSaldoById(inputModel *model.UserSaldo) (*model.UserSaldo, error) {

	sqlScript := `SELECT no_rekening, saldo 
				FROM user_saldo 
				WHERE no_rekening = ?`

	logger.Info("Executing GetUserSaldoById SQL query", "input", *inputModel)

	res := r.db.Raw(sqlScript, inputModel.NoRekening).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("Failed to get user saldo by no_rekening in GetUserSaldoById", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("Successfully retrieved user saldo by no_rekening", "output", *inputModel)

	return inputModel, nil
}

func (r *userSaldoRepository) UpdateUserSaldo(inputModel *model.UserSaldo) (*model.UserSaldo, error) {

	sqlScript := `UPDATE user_saldo
				SET updated_at = ?, saldo = ?  
				WHERE no_rekening = ? 
				RETURNING *`

	logger.Info("Executing UpdateUserSaldo SQL query", "input", *inputModel)

	res := r.db.Raw(sqlScript, time.Now(), inputModel.Saldo, inputModel.NoRekening).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("Failed to update user saldo in UpdateUserSaldo", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("Successfully updated user saldo", "output", *inputModel)

	return inputModel, nil
}
