package main

import (
	"github.com/Barsu5489/commerce/routes"
	"github.com/Barsu5489/commerce/models"
	"net/http"
)

func main() {
	routes.RouteSetup()
	models.SeedProducts() 
	http.ListenAndServe(":8080", nil)
}
