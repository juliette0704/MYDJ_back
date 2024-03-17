package entity

import "time"

type Shot struct {
	ID           uint    `gorm:"primaryKey" json:"id,omitempty"`
	Name         string  `json:"name"`
	Percentage   float32 `json:"percentage"`
	AlreadyTake  bool    `json:"check"`
	Price        float32 `json:"price"`
	Points       int     `json:"points"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
