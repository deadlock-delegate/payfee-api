package models

import (
	"time"

	"github.com/google/uuid"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Link model
type Link struct {
	Base
	Title          string `json:"title" gorm:"size:125"`
	Description    string `json:"description" gorm:"size:512"`
	Fee            uint64 `json:"fee,string" gorm:"size:20`
	PaymentAddress string `json:"payment_address" gorm:"size:34`
	URL            string `json:"url" gorm:"size:512"`
}
