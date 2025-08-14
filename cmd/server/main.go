package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/ricardolindner/go-expert-cloud-run/internal/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file: ", err)
	}

	http.HandleFunc("/weather", handlers.GetWeather)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
