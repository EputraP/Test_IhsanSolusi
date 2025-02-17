package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	NoRekening string         `json:"no_rekening" gorm:"type:varchar;not null"`
	Nama       string         `json:"nama" gorm:"type:varchar;not null"`
	NIK        string         `json:"nik" gorm:"type:varchar;not null"`
	NoHP       string         `json:"no_hp" gorm:"type:varchar;not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
