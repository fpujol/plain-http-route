package main

import (
	"net/http"

	"github.com/fpujol/plain-http-route/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
