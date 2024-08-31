package tea

import (
	"time"

	"github.com/google/uuid"
)

type Tea struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string
	Type      string
	Elevation float64
	Year      int
	Cultivar  string
	StorageID uuid.UUID `gorm:"type:uuid"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
