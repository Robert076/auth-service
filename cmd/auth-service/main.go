package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Robert076/auth-service/internal/constants"
	db_config "github.com/Robert076/auth-service/internal/db/db-config"
	postgres_repository "github.com/Robert076/auth-service/internal/db/repository/postgres"
	register_handler "github.com/Robert076/auth-service/internal/handlers/register"
	validation_service "github.com/Robert076/auth-service/internal/service/validation-service"
	user "github.com/Robert076/auth-service/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("%s: Failed loading env file", constants.ServiceName)
		}
	}

	cfg, err := db_config.LoadDBConfig()
	if err != nil {
		log.Fatalf("%s: Failed to load db config: %v", constants.ServiceName, err)
	}

	dbstrategy, err := cfg.Strategy()
	if err != nil {
		log.Fatalf("%s: Failed to load db strategy: %v", constants.ServiceName, err)
	}

	db, err := db_config.InitDB(dbstrategy)
	if err != nil {
		log.Fatalf("%s: Failed to init db: %v", constants.ServiceName, err)
	}

	defer db.Close()

	repo := postgres_repository.NewPostgresRepository(db)

	http.HandleFunc("/register", register_handler.RegisterHandler(repo))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if err := validation_service.IsValidHttpRequest(r, http.MethodGet); err != nil {
			http.Error(w, "Invalid method for request. This endpoint only accepts GET.", http.StatusBadRequest)
			log.Printf("%s: Error validating request for GET (login). The issue might be that this endpoint only accepts GET rs. Error: %v", constants.ServiceName, err)
			return
		}

		var u user.LoginUserDTO

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			log.Printf("%s: Invalid JSON payload: %v", constants.ServiceName, err)
			return
		}

		if err := validation_service.IsValidUserLogin(u); err != nil {
			http.Error(w, "Invalid user received for login endpoint.", http.StatusBadRequest)
			log.Printf("%s: Invalid user received for login endpoint: %v", constants.ServiceName, err)
			return
		}

	})

	if err := http.ListenAndServe(":"+os.Getenv("ENDPOINT_PORT"), nil); err != nil {
		log.Fatalf("%s: error starting http server: %v", constants.ServiceName, err)
	}
}
