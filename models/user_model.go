package models

import "time"

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
