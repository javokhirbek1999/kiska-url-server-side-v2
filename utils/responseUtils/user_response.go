package utils

import (
	"core/models"
	"core/serializers"
)

func ResponseUserRegister(userObj *models.Core_user) serializers.UserSerializer {
	return serializers.UserSerializer{
		Email:       userObj.Email,
		Username:    userObj.Username,
		Password:    userObj.Password,
		DateJoined:  userObj.DateJoined,
		DateUpdated: userObj.DateUpdated,
		LastLogin:   userObj.LastLogin,
		IsStaff:     userObj.IsStaff,
		IsActive:    userObj.IsActive,
		IsSuperuser: userObj.IsSuperuser,
	}
}
