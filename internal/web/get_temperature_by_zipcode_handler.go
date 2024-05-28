package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetTemperatureByCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := mux.Vars(r)["cep"]

	if cep == "" || len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
}
