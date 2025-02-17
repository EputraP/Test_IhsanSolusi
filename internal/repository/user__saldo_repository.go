package repository

import (
	"time"

	"github.com/EputraP/Test_IhsanSolusi/internal/model"
	"gorm.io/gorm"
)

type UserSaldoRepository interface {
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

	sqlScript := `INSERT INTO users (no_rekening , saldo , created_at) 
				VALUES (?,?,?) 
				RETURNING id;`

	res := r.db.Raw(sqlScript, inputModel.NoRekening, inputModel.Saldo, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r *userSaldoRepository) UpdateUserSaldo(inputModel *model.User) (*model.User, error) {

	sqlScript := `INSERT INTO users (no_rekening , nama , nik, no_hp, created_at) 
				VALUES (?,?,?,?,?) 
				RETURNING no_rekening;`

	res := r.db.Raw(sqlScript, inputModel.NoRekening, inputModel.Nama, inputModel.NIK, inputModel.NoHP, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}
