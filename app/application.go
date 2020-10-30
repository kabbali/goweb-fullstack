package app

import "net/http"

//var (
//	router = fiber.New()
//)

func StartApplication() {
	mapUrls()
	//log.Fatal(router.Listen(":8080"))
	http.ListenAndServe(":8080", nil)
}
