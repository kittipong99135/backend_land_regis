package controllers

import (
	"agent_office/src/database"
	"agent_office/src/util"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func CreateLayerEndpoint(c *fiber.Ctx) error {
	layerName := c.FormValue("layer_name")
	formStatus := c.FormValue("status")
	if formStatus != "true" && formStatus != "false" {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : Invalid status")
	}
	status := true
	if formStatus != "true" {
		status = false
	}
	file, err := c.FormFile("kmz_file")
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : Failed to read file")
	}
	if strings.ToLower(filepath.Ext(file.Filename)) != ".kmz" {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : Invalid file type. Only .kmz files are allowed.")
	}

	type MaxId struct {
		Id string `gorm:"column:max_id"`
	}
	maxId := MaxId{}
	if err := database.DB.Raw("SELECT MAX(CAST(layer_id AS UNSIGNED)) AS max_id FROM tbl_layers").Scan(&maxId).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}

	layerId, err := strconv.Atoi(string(maxId.Id))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("[Error] : " + err.Error())
	}
	helper := util.RequireUtil()

	uploadDir := "public/kmz/"
	fileName := "[L-" + helper.UString.AutoIncarmentAccountId(layerId, 7) + "]:" + layerName + ".kmz"

	log.Println(fileName)
	savePath := filepath.Join(uploadDir, fileName)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : Failed to save file")
	}

	lat, lon, err := helper.UFile.GetPosrKMZ(savePath)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : " + err.Error())
	}

	userCookie := c.Cookies("user_data", "")

	user := struct {
		AccountID   string `json:"account_id"`
		Email       string `json:"email"`
		Exp         int64  `json:"exp"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		RoleOffice  string `json:"role_office"`
		RoleWebsite string `json:"role_website"`
	}{}

	if err := json.Unmarshal([]byte(userCookie), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : " + " Failed to unmarshal JSON " + err.Error())
	}

	layer := database.Layer{
		LayerId:   helper.UString.AutoIncarmentAccountId(layerId, 7),
		LayerName: layerName,
		KmzPath:   uploadDir + fileName,
		Status:    status,
		UpdateBy:  user.FirstName + " " + user.LastName,
	}

	if err := database.DB.Create(&layer).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"status":  "fail",
				"message": "[Error] : Layer data is comflit",
			})
		}

		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": "[Error] : internal server error",
		})
	}

	if lat != nil && lon != nil {
		log.Println(*lat)
		log.Println(*lon)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": nil,
	})
}

func GetallLayerEndpoint(c *fiber.Ctx) error {
	layers := []database.Layer{}
	if err := database.DB.Find(&layers).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": layers,
	})
}
func GetOneLayerEndpoint(c *fiber.Ctx) error {
	layerId := c.Params("layer_id")
	layer := database.Layer{}
	if err := database.DB.First(&layer, "layer_id = ?", layerId).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": layer,
	})
}
func EditLayerEndpoint(c *fiber.Ctx) error {
	layerId := c.Params("layer_id")
	layer := database.Layer{}
	if err := database.DB.First(&layer, "layer_id = ?", layerId).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	layerName := c.FormValue("layer_name")
	if layerName == "" {
		layerName = layer.LayerName
	}

	formStatus := c.FormValue("status")
	if formStatus != "true" && formStatus != "false" {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : Invalid status")
	}
	status := true
	switch formStatus {
	case "":
		status = layer.Status
	case "false":
		status = true
	}

	fileUpload := false
	filePath := layer.KmzPath
	kmzStr := strings.Split(filePath, ":")
	absfilePath, err := filepath.Abs(kmzStr[0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : finding absolute path " + filePath)
	}
	file, err := c.FormFile("kmz_file")
	if err != nil {
		fileUpload = true
	} else {
		if strings.ToLower(filepath.Ext(file.Filename)) != ".kmz" {
			return c.Status(fiber.StatusBadRequest).SendString("[Error] : Invalid file type. Only .kmz files are allowed.")
		}
	}

	if !fileUpload {
		if err := os.Remove(absfilePath); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status":  "fail",
				"message": "Failed to delete file " + err.Error(),
			})
		}

		uploadDir := "public/kmz/"
		fileName := "[L-" + layerId + "]:" + layerName + ".kmz"
		savePath := filepath.Join(uploadDir, fileName)
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("[Error] : Failed to save file")
		}
		filePath = uploadDir + fileName
	}

	layer.LayerName = layerName
	layer.Status = status
	layer.KmzPath = filePath

	if err := database.DB.Model(database.Layer{}).Where("layer_id = ?", layerId).Updates(&layer).Error; err != nil {
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
func DeleteLayerEndpoint(c *fiber.Ctx) error {
	layerId := c.Params("layer_id")
	layer := database.Layer{}
	if err := database.DB.First(&layer, "layer_id = ?", layerId).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	kmzStr := strings.Split(layer.KmzPath, ":")
	absfilePath, err := filepath.Abs(kmzStr[0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : finding absolute path " + layer.KmzPath)
	}

	if err := os.Remove(absfilePath); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "Failed to delete file " + err.Error(),
		})
	}

	if err := database.DB.Where("layer_id = ?", layerId).Delete(&layer).Error; err != nil {
		log.Println(err.Error())
		return c.Status(500).JSON(fiber.Map{
			"status":  "fail",
			"message": "internal server error",
		})
	}

	return nil
}
