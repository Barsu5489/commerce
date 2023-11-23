package routes

import (
	"github.com/Barsu5489/commerce/controllers"
	"net/http"
)

func RouteSetup() {
	http.HandleFunc("/products", controllers.GetProducts)
}
