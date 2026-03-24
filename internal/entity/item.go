package entity

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	Id           uuid.UUID `gorm:"type:uuid;primaryKey"`
	RestaurantId uuid.UUID `gorm:"type:uuid;not null;constraint:OnDelete:CASCADE"`
	Name         string    `gorm:"type:varchar(100);not null"`
	Price        float64   `gorm:"type:decimal(10,2);not null"`
	Available    bool      `gorm:"type:boolean;not null;default:false"`
	ImageURL     string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;autoUpdateTime"`
}
