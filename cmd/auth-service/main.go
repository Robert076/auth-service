package main

import (
	"log"
	"net/http"

	validate "github.com/Robert076/auth-service/internal"
)

func main() {
	const serviceName = "AUTH-SERVICE"
	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		if err := validate.IsValidHttpRequest(request, http.MethodPost); err != nil {
			http.Error(writer, "Invalid method for request. This endpoint only accepts POST.", http.StatusBadRequest)
			log.Printf("%s: Error validating request for POST (Register). The issue might be that this endpoint only accepts POST requests. Error: %v", serviceName, err)
		}
	})
}
