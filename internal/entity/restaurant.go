package entity

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Location  string    `gorm:"type:varchar(50);not null"`
	ImageURL  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;autoCreateTime"`
	Items     []Item    `gorm:"not null;foreignKey:RestaurantId"`
}
