package controllers

import "github.com/gofiber/fiber/v2"

func SignInEndpoint(c *fiber.Ctx) error {
	return c.SendString("Hello; Sign in endpoint")
}
