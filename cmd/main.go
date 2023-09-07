package main

import (
	"log"
	"net/http"

	"github.com/CarApp/api/handler"
	"github.com/CarApp/internal/services"
)

func main() {
	strr := services.NewService()
	// register handlers
	handler.HandlerRequests(strr)
	// http Listen 
	log.Fatal(http.ListenAndServe(":9000", nil))
}
