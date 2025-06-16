package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	db_config "github.com/Robert076/auth-service/internal/db/db-config"
	"github.com/Robert076/auth-service/internal/db/repository"
	postgres_repository "github.com/Robert076/auth-service/internal/db/repository/postgres"
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

	var repo repository.IRepository

	repo = postgres_repository.NewPostgresRepository(db)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodPost); err != nil {
			http.Error(w, "Invalid method for r. This endpoint only accepts POST.", http.StatusBadRequest)
			log.Printf("%s: Error validating r for POST (Register). The issue might be that this endpoint only accepts POST rs. Error: %v", serviceName, err)
			return
		}

		var u user.RegisterUserDTO

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			log.Printf("%s: Invalid JSON payload: %v", serviceName, err)
			return
		}

		if err := validation_service.IsValidUser(u); err != nil {
			http.Error(w, "Invalid user received for register endpoint.", http.StatusBadRequest)
			log.Printf("%s: Invalid user received for register endpoint: %v", serviceName, err)
			return
		}

		if err := repo.RegisterUser(u); err != nil {
			http.Error(w, "Could not insert user in db", http.StatusBadRequest)
			log.Printf("%s: Could not insert user in db: %v", serviceName, err)
			return
		}

		log.Printf("%s: Successfully added user in db.", serviceName)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodGet); err != nil {
			http.Error(w, "Invalid method for r. This endpoint only accepts GET.", http.StatusBadRequest)
			log.Printf("%s: Error validating r for GET (login). The issue might be that this endpoint only accepts GET rs. Error: %v", serviceName, err)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			log.Printf("%s: Invalid JSON payload: %v", serviceName, err)
			return
		}

		if err := validation_service.IsValidUser(u); err != nil {
			http.Error(w, "Invalid user received for login endpoint.", http.StatusBadRequest)
			log.Printf("%s: Invalid user received for login endpoint: %v", serviceName, err)
			return
		}

	})

	if err := http.ListenAndServe(":"+os.Getenv("ENDPOINT_PORT"), nil); err != nil {
		log.Fatalf("%s: error starting http server: %v", serviceName, err)
	}
}
