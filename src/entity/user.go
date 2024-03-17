package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uint `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UUID      uuid.UUID `json:"uuid"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	Password  string    `json:"password,omitempty"`
	UID       string    `json:"uid" gorm:"uniqueIndex,default:null"`
	Shots     []Shot    `gorm:"many2many:user_shots;" json:"shots,omitempty"`
}
