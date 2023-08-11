package models

import (
	"gorm.io/gorm"
	"time"
)

type Choice struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Choice    string    `gorm:"type:varchar(191);not null" json:"choice"`
	Poll      Poll      `gorm:"foreignKey:PollID" json:"poll"`
	PollID    uint64    `json:"poll_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (choice *Choice) BeforeCreate(tx *gorm.DB) (err error) {
	choice.CreatedAt = time.Now()
	choice.UpdatedAt = time.Now()
	return nil
}

func (choice *Choice) BeforeUpdate(tx *gorm.DB) (err error) {
	choice.UpdatedAt = time.Now()
	return nil
}
