package app

import "goweb-fullstack/controllers"

func mapUrls() {
	router.Get("/ping", controllers.Ping)
	router.Get("/", controllers.Hello)
}
