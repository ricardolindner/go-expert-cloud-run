package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/ricardolindner/go-expert-cloud-run/internal/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	http.HandleFunc("/weather", handlers.GetWeather)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
