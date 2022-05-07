package controllers

import (
	"core/database"
	"core/models"
	utils "core/utils"
	responseUtils "core/utils/responseUtils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome")
}

func RegisterUser(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(err.Error())
	}

	var user models.Core_user

	database.Database.DB.Where("email = ?", data["email"]).Find(&user)

	if user.ID != 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User with this email already exists",
		})
	}

	database.Database.DB.Where("user_name = ?", data["user_name"]).Find(&user)

	if user.ID != 0 {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User with this username already exists",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(err.Error())
	}

	user = models.Core_user{
		Email:       data["email"],
		Username:    data["user_name"],
		Password:    password,
		DateJoined:  time.Now(),
		DateUpdated: time.Now(),
		LastLogin:   time.Now(),
		IsStaff:     true,
		IsActive:    false,
		IsSuperuser: false,
	}

	database.Database.DB.Create(&user)

	message := fmt.Sprintf("User %v has been created successfully", user.Username)

	if err := utils.SendMail(user.Email, message); err != nil {
		c.SendStatus(400)
		return c.JSON(err.Error())
	}

	response := responseUtils.ResponseUserRegister(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User has been registered successfully",
		"details": response,
	})
}
