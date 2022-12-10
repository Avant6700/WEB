package model

import "github.com/google/uuid"

type Comics struct {
	UUID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string
	Price       uint64
	Description string
	Year        uint64
	Image       string
}
