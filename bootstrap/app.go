package bootstrap

import (
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

func Start() {
	appEnv := os.Getenv("APP_PORT")
	http.ListenAndServe(":"+appEnv, nil)
}
