package main

import (
	"net/http"

	"github.com/chaudhryjunaid/go-webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
