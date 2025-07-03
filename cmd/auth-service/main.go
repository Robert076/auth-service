package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Robert076/auth-service/internal/constants"
	db_config "github.com/Robert076/auth-service/internal/db/db-config"
	postgres_repository "github.com/Robert076/auth-service/internal/db/repository/postgres"
	authorize_handler "github.com/Robert076/auth-service/internal/handlers/authorize"
	login_handler "github.com/Robert076/auth-service/internal/handlers/login"
	logout_handler "github.com/Robert076/auth-service/internal/handlers/logout"
	register_handler "github.com/Robert076/auth-service/internal/handlers/register"
	"github.com/joho/godotenv"
)

func main() {
	// helloo
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

	repo := postgres_repository.NewPostgresRepository(db) // change this to swap repos (i.e.: mysql_repository.NewMySQLRepository(db))

	http.HandleFunc("/register", register_handler.RegisterHandler(repo))

	http.HandleFunc("/login", login_handler.LoginHandler(repo))

	http.HandleFunc("/authorize", authorize_handler.AuthorizeHandler(repo))

	http.HandleFunc("/logout", logout_handler.LogoutHandler(repo))

	if err := http.ListenAndServe(":"+os.Getenv("ENDPOINT_PORT"), nil); err != nil {
		log.Fatalf("%s: error starting http server: %v", constants.ServiceName, err)
	}
}
