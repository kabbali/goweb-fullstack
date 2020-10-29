package app

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

var (
	router = fiber.New()
)

func StartApplication() {
	mapUrls()
	log.Fatal(router.Listen(":8080"))
}
