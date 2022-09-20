package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Godot struct {
	k string
}

func (ke *Godot) GoDotEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
