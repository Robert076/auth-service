package main

import (
	"log"
	"net/http"

	db_config "github.com/Robert076/auth-service/internal/db/db-config"
	service "github.com/Robert076/auth-service/internal/service"
)

func main() {
	const serviceName = "AUTH-SERVICE"

	db, err := db_config.InitDB()

	if err != nil {
		log.Fatalf("%s: Failed to initialize db", serviceName)
	}

	defer db.Close()

	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		if err := service.IsValidHttpRequest(request, http.MethodPost); err != nil {
			http.Error(writer, "Invalid method for request. This endpoint only accepts POST.", http.StatusBadRequest)
			log.Printf("%s: Error validating request for POST (Register). The issue might be that this endpoint only accepts POST requests. Error: %v", serviceName, err)
		}
	})
}
