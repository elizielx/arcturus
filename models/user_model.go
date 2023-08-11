package models

import (
	"gorm.io/gorm"
	"time"
)

type RoleLevels string

const (
	USER  RoleLevels = "user"
	ADMIN RoleLevels = "admin"
)

type User struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string     `gorm:"type:varchar(191);not null" json:"username"`
	Password   string     `gorm:"type:varchar(191);not null" json:"password"`
	Role       RoleLevels `gorm:"type:role_level" json:"role"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Division   Division   `gorm:"foreignKey:DivisionID" json:"division"`
	DivisionID uint64     `json:"division_id"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return nil
}
