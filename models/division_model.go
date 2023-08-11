package models

import (
	"gorm.io/gorm"
	"time"
)

type Division struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(191);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []User    `gorm:"foreignKey:DivisionID" json:"users"`
}

func (division *Division) BeforeCreate(tx *gorm.DB) (err error) {
	division.CreatedAt = time.Now()
	division.UpdatedAt = time.Now()
	return nil
}

func (division *Division) BeforeUpdate(tx *gorm.DB) (err error) {
	division.UpdatedAt = time.Now()
	return nil
}
