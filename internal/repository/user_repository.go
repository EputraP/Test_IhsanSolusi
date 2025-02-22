package repository

import (
	"time"

	"github.com/EputraP/Test_IhsanSolusi/internal/model"
	"github.com/EputraP/Test_IhsanSolusi/internal/util/logger"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(inputModel *model.User) (*model.User, error)
	CheckUserByNoHPOrNIK(inputModel *model.User) (*model.User, error)
	CheckUserByNoRek(inputModel *model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(inputModel *model.User) (*model.User, error) {

	sqlScript := `INSERT INTO users (no_rekening , nama , nik, no_hp, created_at) 
				VALUES (?,?,?,?,?) 
				RETURNING no_rekening;`

	logger.Info("Executing CreateUser SQL query", "input", *inputModel)

	res := r.db.Raw(sqlScript, inputModel.NoRekening, inputModel.Nama, inputModel.NIK, inputModel.NoHP, time.Now()).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("Failed to create user in CreateUser", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("Successfully created user in CreateUser", "output", *inputModel)

	return inputModel, nil
}

func (r *userRepository) CheckUserByNoHPOrNIK(inputModel *model.User) (*model.User, error) {

	var output *model.User

	sqlScript := `SELECT 
					nik,
					no_hp
				FROM users 
				WHERE 
					(nik= ? OR no_hp = ?) 
					AND deleted_at IS NULL`

	logger.Info("Executing CheckUserByNoHPOrNIK SQL query", "input", *inputModel)

	res := r.db.Raw(sqlScript, inputModel.NIK, inputModel.NoHP).Scan(&output)

	if res.Error != nil {
		logger.Error("Failed to check user by NIK or NoHP in CheckUserByNoHPOrNIK", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("Successfully found user by NIK or NoHP in CheckUserByNoHPOrNIK")

	return output, nil
}

func (r *userRepository) CheckUserByNoRek(inputModel *model.User) (*model.User, error) {

	var output *model.User

	sqlScript := `SELECT 
					no_rekening
				FROM users 
				WHERE 
					no_rekening= ? 
					AND deleted_at IS NULL`

	logger.Info("Executing CheckUserByNoRek SQL query", "input", *inputModel)

	res := r.db.Raw(sqlScript, inputModel.NoRekening).Scan(&output)

	if res.Error != nil {
		logger.Error("Failed to check user by noRek in CheckUserByNoRek", "error", res.Error)
		return nil, res.Error
	}

	logger.Info("Successfully found user by noRek in CheckUserByNoRek", "output", *output)

	return output, nil
}
