package models

import "time"

type Core_user struct {
	ID          uint      `gorm:"primaryKey"`
	Email       string    `gorm:"unique"`
	Username    string    `gorm:"column:user_name" gorm:"unique"`
	Password    []byte    `gorm:"column:password"`
	DateJoined  time.Time `gorm:"column:date_joined"`
	DateUpdated time.Time `gorm:"column:date_updated"`
	LastLogin   time.Time `gorm:"column:last_login"`
	IsStaff     bool      `gorm:"column:is_staff"`
	IsActive    bool      `gorm:"column:is_active"`
	IsSuperuser bool      `gorm:"column:is_superuser"`
}
