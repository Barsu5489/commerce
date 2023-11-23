package controllers

import (
	"github.com/Barsu5489/commerce/models"
	"fmt"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    products := models.GetProducts()
    for _, product := range products {
        fmt.Fprintf(w, "%+v\n", product)
    }
}
