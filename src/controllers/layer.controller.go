package controllers

import (
	"agent_office/src/util"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateLayerEndpoint(c *fiber.Ctx) error {

	layerName := c.FormValue("layer_name")
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : Failed to read file")
	}
	if strings.ToLower(filepath.Ext(file.Filename)) != ".kmz" {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : Invalid file type. Only .kmz files are allowed.")
	}

	uploadDir := "public/kmz/"
	fileName := "[L-" + "0000002]:" + layerName + ".kmz"
	savePath := filepath.Join(uploadDir, fileName)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("[Error] : Failed to save file")
	}

	helper := util.RequireUtil()
	lat, lon, err := helper.UFile.GetPosrKMZ(savePath)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("[Error] : " + err.Error())
	}
	_ = lat
	_ = lon

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": nil,
	})
}

func GetallLayerEndpoint(c *fiber.Ctx) error {
	return nil
}
func GetOneLayerEndpoint(c *fiber.Ctx) error {
	return nil
}
func EditLayerEndpoint(c *fiber.Ctx) error {
	return nil
}
func DeleteLayerEndpoint(c *fiber.Ctx) error {
	return nil
}
