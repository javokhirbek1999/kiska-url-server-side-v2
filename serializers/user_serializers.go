package serializers

import "time"

type UserSerializer struct {
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Password    []byte    `json:"password"`
	DateJoined  time.Time `json:"date_joined"`
	DateUpdated time.Time `json:"date_updated"`
	LastLogin   time.Time `json:"last_login"`
	IsStaff     bool      `json:"is_staff"`
	IsActive    bool      `json:"is_active"`
	IsSuperuser bool      `json:"is_superuser"`
}

type EmailConfirmationSerializer struct {
	Token string `json:"token"`
}
