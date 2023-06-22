package shared

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        string `gorm:"primaryKey;default=gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
