package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Can not load .env")
	}

	API_KEY := os.Getenv("IDOSELL_API_KEY")
	log.Println(API_KEY)
}
