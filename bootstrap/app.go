package bootstrap

import (
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Init()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
}

func Start( router chi.Router) {
	appEnv := os.Getenv("APP_PORT")

	_, err := DatabaseConnection()

	if err != nil {
		log.Fatal(err)
	}


	http.ListenAndServe(":"+appEnv, router)
}

