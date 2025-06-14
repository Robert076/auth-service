package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	db_config "github.com/Robert076/auth-service/internal/db/db-config"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
	user "github.com/Robert076/auth-service/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	const serviceName = "AUTH-SERVICE"

	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("%s: Failed loading env file", serviceName)
		}
	}

	cfg, err := db_config.LoadDBConfig()
	if err != nil {
		log.Fatalf("%s: Failed to load db config: %v", serviceName, err)
	}

	dbstrategy, err := cfg.Strategy()
	if err != nil {
		log.Fatalf("%s: Failed to load db strategy: %v", serviceName, err)
	}

	db, err := db_config.InitDB(dbstrategy)
	if err != nil {
		log.Fatalf("%s: Failed to init db: %v", serviceName, err)
	}

	defer db.Close()

	http.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		if err := validation_service.IsValidHttpRequest(request, http.MethodPost); err != nil {
			http.Error(writer, "Invalid method for request. This endpoint only accepts POST.", http.StatusBadRequest)
			log.Printf("%s: Error validating request for POST (Register). The issue might be that this endpoint only accepts POST requests. Error: %v", serviceName, err)
			return
		}

		var newUser user.User

		if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
			http.Error(writer, "Invalid JSON payload", http.StatusBadRequest)
			log.Printf("%s: Invalid JSON payload: %v", serviceName, err)
			return
		}

		if err := validation_service.IsValidUser(newUser); err != nil {
			http.Error(writer, "Invalid user received for register endpoint.", http.StatusBadRequest)
			log.Printf("%s: Invalid user received for register endpoint: %v", serviceName, err)
			return
		}

		log.Print("User: ", newUser)
	})

	if err := http.ListenAndServe(":"+os.Getenv("ENDPOINT_PORT"), nil); err != nil {
		log.Fatalf("%s: error starting http server: %v", serviceName, err)
	}
}
