package repository

import (
	"gorm.io/gorm"
)

type TestRepository interface {
	TestRepo() (string, error)
}

type testRepository struct {
	db *gorm.DB
}

func NewTestRepository(db *gorm.DB) TestRepository {
	return &testRepository{
		db: db,
	}
}

func (r testRepository) TestRepo() (string, error) {

	return "test repo", nil
}
