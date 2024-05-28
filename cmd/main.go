package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/web"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{zipcode}/temperature", web.GetTemperatureByCepHandler)
	http.ListenAndServe(":8080", r)
}
