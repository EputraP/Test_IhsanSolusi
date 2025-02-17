package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSaldo struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	NoRekening string         `json:"no_rekening" gorm:"type:varchar;not null"`
	Saldo      string         `json:"saldo" gorm:"type:varchar;not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
