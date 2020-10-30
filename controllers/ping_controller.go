package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	//log.Fatal(c.SendStatus(http.StatusOK))
	return c.SendString("pong")
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
