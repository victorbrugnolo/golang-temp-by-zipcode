package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %s", string(err.Error()))
	}

	r := mux.NewRouter()

	r.HandleFunc("/{zipcode}/temperature", web.GetTemperatureByCepHandler)
	http.ListenAndServe(":8080", r)
}
