package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	error := godotenv.Load()

	if error != nil {
		log.Fatal("Erro ao ler o arquivo .env")
	}
}
