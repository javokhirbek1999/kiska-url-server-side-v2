package controllers

import (
	"core/database"
	"core/models"
	hashing "core/utils/hashing"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func ShortenURL(c *fiber.Ctx) error {

	if err := godotenv.Load(".env"); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Internal Server error, it is on us, we will fix it shortly, please try again",
		})
	}

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
			"token":   token,
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

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	var url models.Core_originalurl

	database.Database.DB.Where("url = ?", data["url"]).Where("user_id = ?", user.ID).First(&url)

	// If already exists in Database, do not shorten it, just return it
	if url.ID != 0 {
		// update shortened count
		url.Shortened++
		database.Database.DB.Save(&url)
		return c.Status(200).JSON(fiber.Map{
			"message": "succes",
			"details": url,
		})
	}

	urlHash, err := hashing.HashURL(data["url"], user.ID)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	url = models.Core_originalurl{
		URL:         data["url"],
		DateCreated: time.Now(),
		UserID:      int(user.ID),
		User:        user,
		Shortened:   1,
		UrlHash:     urlHash,
		ShortURL:    c.BaseURL() + "/" + urlHash,
		Visited:     0,
	}

	if err := database.Database.DB.Create(&url).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"details": url,
	})

}

func Redirect(c *fiber.Ctx) error {

	urlHash := c.Params("urlHash")

	var url models.Core_originalurl

	if err := database.Database.DB.Where("urlHash = ?", urlHash).First(&url).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid URL",
		})
	}

	url.Visited++

	database.Database.DB.Save(&url)

	return c.Redirect(url.URL)
}
