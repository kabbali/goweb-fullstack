// Package app provides entry point
package app

import (
	"goweb-fullstack/controllers"
	"net/http"
)

func mapUrls() {
	//router.Get("/ping", controllers.Ping)
	//router.Get("/", controllers.Hello)
	http.HandleFunc("/", controllers.Ping)
}
