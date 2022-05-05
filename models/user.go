package models

import "time"

type Core_user struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Email       string    `json:"email" gorm:"unique"`
	Username    string    `json:"user_name" gorm:"unique"`
	Password    string    `json:"password"`
	DateJoined  time.Time `json:"date_joined"`
	DateUpdated time.Time `json:"date_updated"`
	LastLogin   time.Time `json:"last_login"`
	IsStaff     bool      `json:"is_staff"`
	IsActive    bool      `json:"is_active"`
	IsSuperuser bool      `json:"is_superuser"`
}
