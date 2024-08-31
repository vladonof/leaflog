package tasting

import (
	"time"

	"github.com/google/uuid"
)

type TastingSession struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string
	Date      time.Time
	UserID    uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TastingNote model
type TastingNote struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	SessionID uuid.UUID `gorm:"type:uuid"`
	TeaID     uuid.UUID `gorm:"type:uuid"`
	Notes     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
