package controllers

import (
	"agent_office/src/database"
	"agent_office/src/util"
	"errors"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SignBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignInEndpoint(c *fiber.Ctx) error {
	helper := util.RequireUtil() // [require] : utils helper
	body := SignBody{}

	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	account := database.Account{}
	if err := database.DB.First(&account, "username = ?", body.Username).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(401).JSON(fiber.Map{
				"status":  "fail",
				"message": "username not found",
			})
		}
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	if err := helper.UHash.CompareString(account.Password, body.Password); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"status":  "fail",
			"message": "wrong password",
		})
	}

	payload := map[string]interface{}{
		"account_id":   account.AccountId,
		"first_name":   account.FirstName,
		"last_name":    account.LastName,
		"email":        account.Email,
		"role_office":  account.RoleOfficeId,
		"role_website": account.RoleWebsiteId,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	secretKey := os.Getenv("SECRET_KEY")

	accessToken, err := helper.UJwt.SignToken(payload, secretKey)
	if err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    *accessToken,
		Secure:   false,
		SameSite: "None",
	})

	type Resp struct {
		AccessToken  *string `json:"access_token"`
		RefreshToken *string `json:"refresh_token"`
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": Resp{
			AccessToken:  accessToken,
			RefreshToken: nil,
		},
	})
}
