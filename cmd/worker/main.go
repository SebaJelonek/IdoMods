package main

import (
	"encoding/json"
	"log"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var id int
		id++

		json.NewEncoder(w).Encode(map[string]interface{}{
			"Status":  "Success",
			"Message": "Root",
			"id":      id,
		})
	})

	PORT := os.Getenv("PORT")
	log.Printf("starting server at port %s", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
