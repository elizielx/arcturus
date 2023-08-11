package models

import (
	"gorm.io/gorm"
	"time"
)

type Vote struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Choice     Choice    `gorm:"foreignKey:ChoiceID" json:"choice"`
	ChoiceID   uint64    `json:"choice_id"`
	User       User      `gorm:"foreignKey:UserID" json:"user"`
	UserID     uint64    `json:"user_id"`
	Poll       Poll      `gorm:"foreignKey:PollID" json:"poll"`
	PollID     uint64    `json:"poll_id"`
	Division   Division  `gorm:"foreignKey:DivisionID" json:"division"`
	DivisionID uint64    `json:"division_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (vote *Vote) BeforeCreate(tx *gorm.DB) (err error) {
	vote.CreatedAt = time.Now()
	vote.UpdatedAt = time.Now()
	return nil
}

func (vote *Vote) BeforeUpdate(tx *gorm.DB) (err error) {
	vote.UpdatedAt = time.Now()
	return nil
}
