package main

import (
	"log"
	"net/http"

	"github.com/CarApp/api/handler"
	"github.com/CarApp/internal/services"
)

func main() {
	// register handlers
	strr := services.NewService()
	handler.HandlerRequests(strr)
	log.Fatal(http.ListenAndServe(":9000", nil))
}


/* Pending
1. Logs
2. Tests
*/
