package main

import (
	"net/http"

	"github.com/imarchanka/webservice-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
