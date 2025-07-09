package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/joho/godotenv"
)

type Body struct {
	Params struct {
		OrdersStatuses []string `json:"orderStatuses"`
	} `json:"params"`
}

func main() {
	var id int = 1
	err := godotenv.Load()
	if err != nil {
		log.Panic("Can not load .env")
	}

	API_KEY := os.Getenv("IDOSELL_API_KEY")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id++
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Status":  "Success",
			"Message": "Root",
			"id":      id,
		})
	})
	var payload Body
	payload.Params.OrdersStatuses = []string{
		"new",
		"finished",
		"false",
		"lost",
		"on_order",
		"packed",
		"ready",
		"canceled",
		"payment_waiting",
		"delivery_waiting",
		"suspended",
		"joined",
		"finished_ext",
	}

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(&payload); err != nil {
		log.Fatalf("encode body %v", err)
	}

	req, err := http.NewRequest("POST", "https://zooart6.yourtechnicaldomain.com/api/admin/v5/orders/orders/search", buf)
	if err != nil {
		log.Panicln("request error")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", API_KEY)

	raw, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatalf("dump request, %v", err)
	}
	log.Printf("\n--- REQUEST ---\n%s\n---------------", raw)

	PORT := os.Getenv("PORT")
	log.Printf("starting server at port %s", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
