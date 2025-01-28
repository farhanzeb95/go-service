package main

import (
	"go-web-service/models/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
