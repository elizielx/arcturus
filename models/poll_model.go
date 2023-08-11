package models

import "time"

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
