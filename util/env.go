package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, found := os.LookupEnv("ENV_FILE"); !found {
		os.Setenv("ENV_FILE", ".env")
	}

	envfile := os.Getenv("ENV_FILE")

	if err := godotenv.Load(envfile); err != nil {
		log.Println(err.Error())
	}
}
