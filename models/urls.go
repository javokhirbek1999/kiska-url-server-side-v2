package models

import "time"

type Core_originalurl struct {
	ID          uint      `gorm:"column:id" gorm:"primaryKey"`
	URL         string    `gorm:"colum:url"`
	DateCreated time.Time `gorm:"column:date_created"`
	UserID      int       `gorm:"column:user_id"`
	User        Core_user `gorm:"foreignKey:UserID"`
	Shortened   uint      `gorm:"column:shortened"`
	UrlHash     string    `gorm:"column:urlHash"`
	ShortURL    string    `gorm:"column:shortURL"`
	Visited     uint      `gorm:"column:visited"`
}
