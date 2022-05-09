package serializers

import (
	"core/models"
	"time"
)

type UrlSerializer struct {
	ID          uint      `json:"id"`
	URL         string    `json:"url"`
	DateCreated time.Time `json:"date_created"`
	UserID      int       `json:"user_id"`
	User        models.Core_user
	Shortened   uint   `json:"shortened"`
	UrlHash     string `json:"urlHash"`
	ShortURL    string `json:"shortURL"`
	Visited     uint   `json:"visited"`
}
