package controllers

import (
	"fmt"
	"net/http"
)

//import (
//	"github.com/gofiber/fiber/v2"
//)
//
//func Ping(c *fiber.Ctx) error {
//	//log.Fatal(c.SendStatus(http.StatusOK))
//	return c.SendString("pong")
//}
//
//func Hello(c *fiber.Ctx) error {
//	return c.SendString("Hello, World 👋!")
//}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Pong</h1>")
	fmt.Fprint(w, "<p>Hello, Welcome to my website 👋!</p>")
}