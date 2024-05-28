package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorbrugnolo/golang-temp-cep/internal/web"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{cep}/temperature", web.GetTemperatureByCepHandler)
	http.ListenAndServe(":8080", r)
}
