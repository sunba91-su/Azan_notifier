package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func GetEnv(env string) string {
	return os.Getenv(env)
}

func GetCityInfo()
