package controllers

import (
	"agent_office/src/database"
	"agent_office/src/models"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type createdPermission struct {
	PermissionName string `json:"permission_name"`
	Module         string `json:"module"`
}

// CreatePermissionEndpoint godoc
//
//	@Summary		Handler func create permission endpoint
//	@Description	create new permission endpoint
//	@Tags			Permission
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createdPermission	true	"Request body for created role"
//	@Success		200		{object}	string	"data"
//	@Failure		401		{object}	string				"Unauthorized error"
//	@Failure		404		{object}	string				"Data not found error"
//	@Failure		500		{object}	string				"Internal server error"
//	@Router			/permission/ [post]
func CreatePermissionEndpoint(c *fiber.Ctx) error {
	body := createdPermission{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	permission := database.Permission{}
	nameStr := strings.Split(body.PermissionName, " ")
	permission.PermissionId = "PERM"
	for _, str := range nameStr {
		permission.PermissionId += "_" + strings.ToUpper(str)
	}
	permission.PermissionName = body.PermissionName
	permission.Module = body.Module

	if err := database.DB.Create(&permission).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("creates success")
}

// GetallPermissionEndpoint godoc
//
//	@Summary		Get all permission
//	@Description	get GetallPermissionEndpoint
//	@Tags			Permission
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.Permission	"data"
//	@Failure		401	{object}	string				"Unauthorized error"
//	@Failure		404	{object}	string				"Data not found error"
//	@Failure		500	{object}	string				"Internal server error"
//	@Router			/permission/ [get]
func GetallPermissionEndpoint(c *fiber.Ctx) error {
	permissions := []database.Permission{}

	if err := database.DB.Find(&permissions).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": permissions,
	})
}

// GetPermissionEndpoint godoc
//
//	@Summary		Get  one permission
//	@Description	get GetPermissionEndpoint
//	@Tags			Permission
//	@Accept			json
//	@Produce		json
//	@Param			permission_id	path		string				true	"Pemission ID"
//	@Success		200				{object}	models.Permission	"data"
//	@Failure		401				{object}	string				"Unauthorized error"
//	@Failure		404				{object}	string				"Data not found error"
//	@Failure		500				{object}	string				"Internal server error"
//	@Router			/permission/{permission_id} [get]
func GetPermissionEndpoint(c *fiber.Ctx) error {
	permissionId := c.Params("permission_id")

	log.Println(permissionId)

	permission := database.Permission{}
	if err := database.DB.First(&permission, "permission_id  = ?", permissionId).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": permission,
	})
}

// EditPermissionEndpoint godoc
//
//	@Summary		Edit permission
//	@Description	put EditPermissionEndpoint
//	@Tags			Permission
//	@Accept			json
//	@Produce		json
//	@Param			permission_id	path		string				true	"Pemission ID"
//	@Param			request			body		models.Role			true	"Request body for created role"
//	@Success		200				{object}	models.Permission	"data"
//	@Failure		401				{object}	string				"Unauthorized error"
//	@Failure		404				{object}	string				"Data not found error"
//	@Failure		500				{object}	string				"Internal server error"
//	@Router			/permission/{permission_id} [put]
func EditPermissionEndpoint(c *fiber.Ctx) error {
	permissionId := c.Params("permission_id")

	permission := models.Permission{}
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	now := time.Now()
	if err := database.DB.Model(models.Permission{}).Where("permission_id = ?", permissionId).Updates(&models.Permission{
		Module:    permission.Module,
		UpdatedAt: &now,
	}).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": nil,
	})
}

// DeletePermissionEndpoint godoc
//
//	@Summary		Delete permission
//	@Description	delete DeletePermissionEndpoint
//	@Tags			Permission
//	@Accept			json
//	@Produce		json
//	@Param			permission_id	path		string				true	"Pemission ID"
//	@Success		200				{object}	models.Permission	"data"
//	@Failure		401				{object}	string				"Unauthorized error"
//	@Failure		404				{object}	string				"Data not found error"
//	@Failure		500				{object}	string				"Internal server error"
//	@Router			/permission/{permission_id} [delete]
func DeletePermissionEndpoint(c *fiber.Ctx) error {
	permissionId := c.Params("permission_id")

	permission := models.Permission{}
	if err := database.DB.Where("permission_id = ?", permissionId).Delete(&permission).Error; err != nil {
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
