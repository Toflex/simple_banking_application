package models

import (
	"gorm.io/gorm"
	"simple_banking_application/internal/common/utils"
	"time"
)

type Base struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = utils.GenerateUUID()
	}

	return
}
