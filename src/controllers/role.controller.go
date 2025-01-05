package controllers

import (
	"agent_office/src/database"
	"agent_office/src/models"
	"agent_office/src/util"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// [module] : role
type roleCreated struct {
	RoleName string `json:"role_name"`
	RoleRef  string `json:"role_ref"`
}

// CreateRoleEndpoint godoc
//
//	@Summary		handler func  Create role endpoint
//	@Description	create  new role by request body
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			request	body		roleCreated	true	"request body"
//	@Success		200		{object}	string		"data"
//	@Failure		401		{object}	string		"Unauthorized error"
//	@Failure		404		{object}	string		"Data not found error"
//	@Failure		500		{object}	string		"Internal server error"
//	@Router			/role/ [post]
func CreateRoleEndpoint(c *fiber.Ctx) error {
	body := roleCreated{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	role := database.Role{}
	role.RoleId = "ROLE_" + strings.ToUpper(body.RoleName)
	role.RoleName = body.RoleName
	role.RoleRef = body.RoleRef

	if err := database.DB.Create(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": nil,
	})
}

// GetallRoleEndpoint godoc
//
//	@Summary		Get all role
//	@Description	get GetallRoleEndpoint
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.Role	"role data"
//	@Failure		401	{object}	string			"Unauthorized error"
//	@Failure		404	{object}	string			"Data not found error"
//	@Failure		500	{object}	string			"Internal server error"
//	@Router			/role/ [get]
func GetallRoleEndpoint(c *fiber.Ctx) error {
	roles := []database.Role{}

	if err := database.DB.Find(&roles).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": roles,
	})
}

// GetoneRoleEndpoint godoc
//
//	@Summary		Get one role
//	@Description	put GetoneRoleEndpoint
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			role_id	path		string		true	"Role ID"
//	@Success		200		{object}	models.Role	"data"
//	@Failure		401		{object}	string		"Unauthorized error"
//	@Failure		404		{object}	string		"Data not found error"
//	@Failure		500		{object}	string		"Internal server error"
//	@Router			/role/:role_id [get]
func GetoneRoleEndpoint(c *fiber.Ctx) error {
	roleId := c.Params("role_id")

	role := database.Role{}
	if err := database.DB.First(&role, "role_id = ?", roleId).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": role,
	})
}

// EditRolePermissionEndpoint godoc
//
//	@Summary		Edit permission role
//	@Description	put EditRolePermissionEndpoint
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			role_id	path		string					true	"Role ID"
//	@Param			request	body		models.PermissionJson	true	"Request body for editing an role permission"
//	@Success		200		{object}	string					"data"
//	@Failure		401		{object}	string					"Unauthorized error"
//	@Failure		404		{object}	string					"Data not found error"
//	@Failure		500		{object}	string					"Internal server error"
//	@Router			/role/{role_id} [put]
func EditRolePermissionEndpoint(c *fiber.Ctx) error {
	roleId := c.Params("role_id")

	role := models.PermissionJson{}
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	if err := database.DB.Model(models.Role{}).Where("role_id = ?", roleId).Updates(&models.Role{
		Permissions: role,
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

// [module] : role permission

type createdRolePermission struct {
	RoleRefId   string                       `json:"role_id"`
	Permissions []models.GroupRolePermission `json:"group_permission"`
}

// CreateRolePermissionEndpoint godoc
//
//	@Summary		Create role permission
//	@Description	post CreateRolePermissionEndpoint
//	@Tags			RolePermission
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.RolePermissionBody	true	"Request body for created role"
//	@Success		200		{object}	string						"data"
//	@Failure		401		{object}	string						"Unauthorized error"
//	@Failure		404		{object}	string						"Data not found error"
//	@Failure		500		{object}	string						"Internal server error"
//	@Router			/role/role_permission/ [post]
func CreateRolePermissionEndpoint(c *fiber.Ctx) error {
	helper := util.RequireUtil()

	body := createdRolePermission{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	type MaxId struct {
		Id string `gorm:"column:max_id"`
	}
	maxId := MaxId{}
	if err := database.DB.Raw("  SELECT MAX(CAST(REPLACE(role_permission_id, 'RP', '') AS UNSIGNED)) AS max_id FROM tbl_role_permission").Scan(&maxId).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	rolePermissionId, err := strconv.Atoi(string(maxId.Id))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	newRp := make([]database.RolePermission, 0)

	for index, permission := range body.Permissions {
		newRp = append(newRp, database.RolePermission{
			RolePermissionId: "RP" + helper.UString.AutoIncarmentAccountId(rolePermissionId+index, 5),
			RoleRefId:        body.RoleRefId,
			PermissionRefId:  permission.PermissionId,
			Permissions:      permission.Permission,
		})
	}

	if err := database.DB.Create(&newRp).Error; err != nil {
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

// GetRolePermissionEndPoint godoc
//
//	@Summary		Get one role permission
//	@Description	get GetRolePermissionEndPoint
//	@Tags			RolePermission
//	@Accept			json
//	@Produce		json
//	@Param			role_permiassion_id	path		string						true	"Role ID"
//	@Success		200					{object}	models.RolePermissionBody	"data"
//	@Failure		401					{object}	string						"Unauthorized error"
//	@Failure		404					{object}	string						"Data not found error"
//	@Failure		500					{object}	string						"Internal server error"
//	@Router			/role_permission/{role_permiassion_id} [get]
func GetRolePermissionEndPoint(c *fiber.Ctx) error {
	rolePermissionId := c.Params("role_id")
	type queryRes struct {
		RolePermissionId string                `gorm:"primaryKey;column:role_permission_id;unique;size:50" json:"role_permission_id"`
		RoleId           string                `gorm:"column:role_id;size:100" json:"role_id"`
		PermissionId     string                `gorm:"column:permission_id" json:"permission_id"`
		Permissions      models.PermissionJson `gorm:"column:permissions;type:json" json:"permissions"`
		RoleName         string                `gorm:"uniqueIndex;column:role_name;unique;size:100" json:"role_name"`
		PermissionName   string                `gorm:"column:permission_name;size:100" json:"permission_name"`
		Module           string                `gorm:"column:module;size:100" json:"module"`
	}
	rolePermission := []queryRes{}

	sql := `select rp.*, r.role_name, p.permission_name, p.module from tbl_role_permission as rp 
		left join tbl_roles  as r on rp.role_id = r.role_id and r.role_id = ?
		left join tbl_permissions  as p on rp.permission_id = p.permission_id
		where rp.role_id =?`

	if err := database.DB.Raw(sql, rolePermissionId, rolePermissionId).Scan(&rolePermission).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": "[Error] : internal server error",
		})
	}

	respPermission := make([]models.GroupRole, 0)
	for _, permission := range rolePermission {
		respPermission = append(respPermission, models.GroupRole{
			PermissionId:   permission.PermissionId,
			PermissionName: permission.PermissionName,
			Module:         permission.Module,
			Permission:     permission.Permissions,
		})
	}

	resp := models.RolePermissionBody{
		RolePermissionId: rolePermission[0].RolePermissionId,
		RoleId:           rolePermission[0].RoleId,
		RoleName:         rolePermission[0].RoleName,
		PermissionGroup:  respPermission,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": resp,
	})
}

// GetallRolePermissionEndPoint godoc
//
//	@Summary		Get all role permission
//	@Description	get GetallRolePermissionEndPoint
//	@Tags			RolePermission
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]models.RolePermissionBody	"data"
//	@Failure		401	{object}	string						"Unauthorized error"
//	@Failure		404	{object}	string						"Data not found error"
//	@Failure		500	{object}	string						"Internal server error"
//	@Router			/role_permission [get]
func GetallRolePermissionEndPoint(c *fiber.Ctx) error {
	type queryRes struct {
		RolePermissionId string                `gorm:"primaryKey;column:role_permission_id;unique;size:50" json:"role_permission_id"`
		RoleId           string                `gorm:"column:role_id;size:100" json:"role_id"`
		PermissionId     string                `gorm:"column:permission_id" json:"permission_id"`
		Permissions      models.PermissionJson `gorm:"column:permissions;type:json" json:"permissions"`
		RoleName         string                `gorm:"uniqueIndex;column:role_name;unique;size:100" json:"role_name"`
		PermissionName   string                `gorm:"column:permission_name;size:100" json:"permission_name"`
		Module           string                `gorm:"column:module;size:100" json:"module"`
	}
	rolePermission := []queryRes{}

	sql := `select rp.*, r.role_name, p.permission_name, p.module from tbl_role_permission as rp 
	left join tbl_roles as r on rp.role_id = r.role_id 
	left join tbl_permissions as p on p.permission_id = rp.permission_id`

	if err := database.DB.Raw(sql).Scan(&rolePermission).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": "[Error] : internal server error",
		})
	}

	resp := make([]models.RolePermissionBody, 0)
	for _, role := range rolePermission {
		existsRole := false
		existsIndex := -1
		if len(resp) > 0 {
			for index, ele := range resp {
				if ele.RoleId == role.RoleId {
					existsRole = true
					existsIndex = index
				}
			}
		}

		if !existsRole {
			permissionGroup := make([]models.GroupRole, 0)
			permissionGroup = append(permissionGroup, models.GroupRole{
				PermissionId:   role.PermissionId,
				PermissionName: role.PermissionName,
				Module:         role.Module,
				Permission:     role.Permissions,
			})
			resp = append(resp, models.RolePermissionBody{
				RolePermissionId: role.RolePermissionId,
				RoleId:           role.RoleId,
				RoleName:         role.RoleName,
				PermissionGroup:  permissionGroup,
			})
		} else {
			resp[existsIndex].PermissionGroup = append(resp[existsIndex].PermissionGroup, models.GroupRole{
				PermissionId:   role.PermissionId,
				PermissionName: role.PermissionName,
				Module:         role.Module,
				Permission:     role.Permissions,
			})
		}

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": resp,
	})
}

type updateRolePermission struct {
}

// UpdateRolePermissionEndpoint godoc
//
//	@Summary		Update role permission detail
//	@Description	put UpdateRolePermissionEndpoint
//	@Tags			RolePermission
//	@Accept			json
//	@Produce		json
//	@Param			role_id	path		string						true	"Role ID"
//	@Param			request	body		models.RolePermissionBody	true	"Request body for editing an role permission"
//	@Success		200		{object}	string						"data"
//	@Failure		401		{object}	string						"Unauthorized error"
//	@Failure		404		{object}	string						"Data not found error"
//	@Failure		500		{object}	string						"Internal server error"
//	@Router			/role/{role_id}/permission_detail [put]
func UpdateRolePermissionEndpoint(c *fiber.Ctx) error {
	roleId := c.Params("role_id")

	body := models.RolePermissionBody{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}
	for _, permission := range body.PermissionGroup {
		log.Println(permission.PermissionId)
		spew.Dump(permission.Permission)

		if err := database.DB.Model(database.RolePermission{}).
			Where("role_id = ? and permission_id = ?", roleId, permission.PermissionId).
			Select("permissions").
			Updates(&database.RolePermission{
				Permissions: database.JsonPermission{
					Create: permission.Permission.Create,
					View:   permission.Permission.View,
					Edit:   permission.Permission.Edit,
					Delete: permission.Permission.Delete,
				},
			},
			).Error; err != nil {
			log.Println(err.Error())
			return c.Status(500).JSON(fiber.Map{
				"status":  "fail",
				"message": "internal server error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": nil,
	})
}
