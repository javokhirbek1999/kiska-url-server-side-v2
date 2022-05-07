package controllers

import (
	"core/database"
	"core/models"
	utils "core/utils"
	responseUtils "core/utils/responseUtils"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome")
}

func RegisterUser(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"details": err.Error(),
		})
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

	var lastUser models.Core_user

	database.Database.DB.Last(&lastUser)

	userID := 0

	if lastUser.ID == 0 {
		userID = 1
	} else {
		userID = int(lastUser.ID) + 1
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.StandardClaims{
		Issuer:    strconv.Itoa(userID),
		ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
	})

	if err := godotenv.Load(".env"); err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unknown Error, it is on us, we will fix it shortly",
		})
	}

	verification_token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Uknow Error, please try again",
		})
	}

	link := fmt.Sprintf("https://127.0.0.1:8000/api/auth/confirm/%v", verification_token)

	message := fmt.Sprintf("Email verification for %v \nPlease click the link below to verify your email and activate your account\n\nLink: %v\n\nSincerely,\nKiska URL", user.Username, link)

	if err := utils.SendMail(user.Email, message); err != nil {
		fmt.Println("Error")
		c.SendStatus(400)
		return c.JSON(err.Error())
	}

	database.Database.DB.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Verification link has been sent to your email",
	})
}

func EmailConfirmation(c *fiber.Ctx) error {

	if err := godotenv.Load(".env"); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"details": "We have encountered an error while processing your request, we will fix it soon",
		})
	}

	verification_token := c.Params("token")

	token, err := jwt.ParseWithClaims(verification_token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.Core_user

	if err := database.Database.DB.Where("id = ?", claims.Issuer).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	user.IsActive = true

	database.Database.DB.Save(&user)

	response := responseUtils.ResponseUserRegister(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Verification successfully",
		"details": response,
	})

}
