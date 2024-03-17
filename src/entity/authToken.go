package entity

import (
	"gorm.io/gorm"
)

type AuthToken struct {
	gorm.Model
	Token  string `json:"token"`
	UserID uint64 `gorm:"primary_key" gorm:"uniqueIndex"`
	User   User   `gorm:"ForeignKey:UserID" gorm:"uniqueIndex" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
