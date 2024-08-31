package storage

import (
	"time"

	"github.com/google/uuid"
)

type Storage struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string
	Location    string
	Size        float64
	Temperature float64
	Humidity    float64
	TeaTypes    []string  `gorm:"type:text[]"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
