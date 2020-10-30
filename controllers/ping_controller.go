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
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, Welcome to my website!</h1>")
	} else if r.URL.Path == "/ping" {
		fmt.Fprint(w, "<h1>Pong</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please contact me by e-mail" +
			" to <a href=\"mailto:support@kabbali.con\">support@kabbali.com</a>.")
	}


}