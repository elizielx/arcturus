package models

import (
	"gorm.io/gorm"
	"time"
)

type Poll struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(191);not null" json:"name"`
	Description string    `gorm:"type:varchar(191);not null" json:"description"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `gorm:"index" json:"deleted_at"`
	Creator     User      `gorm:"foreignKey:CreatedBy" json:"creator"`
	CreatedBy   uint64    `json:"created_by"`
}

func (poll *Poll) BeforeCreate(tx *gorm.DB) (err error) {
	poll.CreatedAt = time.Now()
	poll.UpdatedAt = time.Now()
	return nil
}

func (poll *Poll) BeforeUpdate(tx *gorm.DB) (err error) {
	poll.UpdatedAt = time.Now()
	return nil
}

func (poll *Poll) BeforeDelete(tx *gorm.DB) (err error) {
	poll.DeletedAt = time.Now()
	return nil
}
