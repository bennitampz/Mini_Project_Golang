package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID"`
}
