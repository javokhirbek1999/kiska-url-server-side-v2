package models

import "time"

type Core_user struct {
	ID          uint      `gorm:"primaryKey"`
	Email       string    `gorm:"unique"`
	Username    string    `gorm:"unique column:user_name"`
	Password    string    `gorm:"column:password"`
	DateJoined  time.Time `gorm:"column:date_joined"`
	DateUpdated time.Time `gorm:"column:date_updated"`
	LastLogin   time.Time `gorm:"column:last_login"`
	IsStaff     bool      `gorm:"column:is_staff"`
	IsActive    bool      `gorm:"column:is_active"`
	IsSuperuser bool      `gorm:"column:is_superuser"`
}
