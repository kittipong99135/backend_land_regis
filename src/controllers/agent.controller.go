package controllers

import (
	"agent_office/src/database"
	"agent_office/src/models"
	"agent_office/src/util"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type accountCreated struct {
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	PhoneNumber     string  `json:"phone_number"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
	OtpCode         string  `json:"otp_code"`
	AzureAdId       *string `json:"azure_ad_id"`
	AuthType        string  `json:"auth_type"`
	FirstName       string  `json:"firstname"`
	LastName        string  `json:"lastname"`
}

// CreateAgentEndpoint godoc
//
//	@Summary		Handler func create agent endpoinrt
//	@Description	create new agent by request body
//	@Tags			Agent
//	@Accept			json
//	@Produce		json
//	@Param			request	body		accountCreated	true	"request body"
//	@Success		200		{object}	string			"data"
//	@Failure		401		{object}	string			"Unauthorized error"
//	@Failure		404		{object}	string			"Data not found error"
//	@Failure		500		{object}	string			"Internal server error"
//	@Router			/agent/ [post]
func CreateAgentEndpoinrt(c *fiber.Ctx) error {
	helper := util.RequireUtil() // [require] : utils helper

	// [process] : reqquest bodt parser
	body := accountCreated{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	// [process] : auto increment
	account := database.Account{}
	type MaxId struct {
		Id string `gorm:"column:max_id"`
	}

	maxId := MaxId{}
	if err := database.DB.Raw("SELECT MAX(CAST(account_id AS UNSIGNED)) AS max_id FROM tbl_accounts").Scan(&maxId).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}
	accountId, err := strconv.Atoi(string(maxId.Id))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	// [process] : password hashed
	if body.Password != body.ConfirmPassword {
		c.Status(fiber.ErrBadRequest.Code).SendString("[Error] : password miss match")
	}
	passwordHash, err := helper.UHash.HashedString(body.Password)
	if err != nil {
		c.Status(fiber.ErrInternalServerError.Code).SendString("[Error] : " + err.Error())
	}

	// [process] : insert new account
	account.AccountId = helper.UString.AutoIncarmentAccountId(accountId, 7)
	account.Password = passwordHash
	account.Username = body.Username
	account.Email = body.Email
	account.PhoneNumber = body.PhoneNumber
	account.OtpCode = body.OtpCode
	account.AzureAdId = body.AzureAdId
	account.AuthType = body.AuthType
	account.FirstName = body.FirstName
	account.LastName = body.LastName
	account.RoleOfficeId = "ROLE_GUST"
	account.RoleWebsiteId = "ROLE_GUST"

	if err := database.DB.Create(&account).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  "fail",
				"message": "[Error] : Email or Username is comflit",
			})
		}

		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": "[Error] : internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": nil,
	})
}

// GetallAgentEndpoint godoc
//
//	@Summary		Handler func get all agent endpoint
//	@Description	get all agent
//	@Tags			Agent
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]database.Account "data"
//	@Failure		401	{object}	string				"Unauthorized error"
//	@Failure		404	{object}	string				"Data not found error"
//	@Failure		500	{object}	string				"Internal server error"
//	@Router			/agent/ [get]
func GetallAgentEndpoint(c *fiber.Ctx) error {
	sql := ""

	roleWeb := c.Query("role_web", "all")
	roleOffice := c.Query("role_office", "all")

	findRoleWeb := []string{}
	if roleWeb != "all" {
		roleWebs := strings.Split(roleWeb, ",")
		findRoleWeb = append(findRoleWeb, roleWebs...)
	}
	for _, role := range findRoleWeb {
		if len(sql) > 0 {
			sql += " or role_website_id = '" + role + "'"
		} else {
			sql += "role_website_id = '" + role + "'"
		}
	}

	findRoleOffice := []string{}
	if roleOffice != "all" {
		roleOffices := strings.Split(roleOffice, ",")
		findRoleOffice = append(findRoleOffice, roleOffices...)
	}

	for _, role := range findRoleOffice {
		if len(sql) > 0 {
			sql += " or role_office_id = '" + role + "'"
		} else {
			sql += "role_office_id = '" + role + "'"
		}
	}

	position := c.Query("position", "all")
	findPositions := []string{}
	if position != "all" {
		positions := strings.Split(position, ",")
		findPositions = append(findPositions, positions...)
	}

	for _, position := range findPositions {
		if len(sql) > 0 {
			sql += " or position = '" + position + "'"
		} else {
			sql += "position = '" + position + "'"
		}
	}

	status := c.Query("status", "all")
	findStatus := []string{}
	if status != "all" {
		arrStatus := strings.Split(status, ",")
		findStatus = append(findStatus, arrStatus...)
	}

	for _, status := range findStatus {
		if len(sql) > 0 {
			sql += " or status = '" + status + "'"
		} else {
			sql += "status = '" + status + "'"
		}
	}

	accounts := []database.Account{}
	if err := database.DB.
		Preload("RoleOffice").
		Preload("RoleWebsite").
		Find(&accounts, sql).
		Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": accounts,
	})
}

// GetOneAgentEndpoint godoc
//
//	@Summary		Handler func get one agent by id
//	@Description	get agent by id
//	@Tags			Agent
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.GetAccount	"data"
//	@Failure		401	{object}	string				"Unauthorized error"
//	@Failure		404	{object}	string				"Data not found error"
//	@Failure		500	{object}	string				"Internal server error"
//	@Router			/agent/:agent_id [get]
func GetOneAgentEndpoint(c *fiber.Ctx) error {
	agentId := c.Params("agent_id")
	account := database.Account{}

	if err := database.DB.
		Preload("RoleOffice").
		Preload("RoleWebsite").
		First(&account, "account_id = ?", agentId).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": account,
	})
}

// EditAgentEndpointntEndpoint godoc
//
//	@Summary		Edit agent
//	@Description	edit EditAgentEndpoint
//	@Tags			Agent
//	@Accept			json
//	@Produce		json
//	@Param			agent_id	path		string				true	"Agent ID"
//	@Param			request		body		models.GetAccount	true	"Request body for editing an agent"
//	@Success		200			{object}	string				"data"
//	@Failure		401			{object}	string				"Unauthorized error"
//	@Failure		404			{object}	string				"Data not found error"
//	@Failure		500			{object}	string				"Internal server error"
//	@Router			/agent/:agent_id [put]
func EditAgentEndpoint(c *fiber.Ctx) error {
	agentId := c.Params("agent_id")

	account := models.GetAccount{}
	if err := c.BodyParser(&account); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	account.AccountID = agentId

	if err := database.DB.Model(models.GetAccount{}).Where("account_id = ?", agentId).Updates(&account).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": nil,
	})
}

type updatedRole struct {
	RoleId string `json:"role_id"`
}

func UpdateRoleAgentEndpoint(c *fiber.Ctx) error {
	agentId := c.Params("agent_id")
	roleRef := c.Query("ref", "")

	role := updatedRole{}
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	switch roleRef {
	case "office":
		if err := database.DB.Model(database.Account{}).Where("account_id = ?", agentId).Updates(&database.Account{
			RoleOfficeId: role.RoleId,
		}).Error; err != nil {
			log.Println(err.Error())
			return c.Status(500).JSON(fiber.Map{
				"status":  "fail",
				"message": "internal server error",
			})
		}
	case "website":
		if err := database.DB.Model(database.Account{}).Where("account_id = ?", agentId).Updates(&database.Account{
			RoleWebsiteId: role.RoleId,
		}).Error; err != nil {
			log.Println(err.Error())
			return c.Status(500).JSON(fiber.Map{
				"status":  "fail",
				"message": "internal server error",
			})
		}
	default:
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "role ref not found",
		})
	}
	return nil
}

// DeleteAgentEndpoint godoc
//
//	@Summary		Delete agent
//	@Description	delete DeleteAgentEndpoint
//	@Tags			Agent
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string	"data"
//	@Failure		401	{object}	string	"Unauthorized error"
//	@Failure		404	{object}	string	"Data not found error"
//	@Failure		500	{object}	string	"Internal server error"
//	@Router			/agent/:agent_id [delete]
func DeleteAgentEndpoint(c *fiber.Ctx) error {
	agentId := c.Params("agent_id")

	account := models.GetAccount{}

	if err := database.DB.Where("account_id = ?", agentId).Delete(&account).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": nil,
	})
}

func DownloadAgentEndpoint(c *fiber.Ctx) error {
	return c.SendString("Hello; Download agent endpoint")
}
