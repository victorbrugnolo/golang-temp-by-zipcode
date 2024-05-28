package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorbrugnolo/golang-temp-zipcode/internal/usecase"
)

func GetTemperatureByCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := mux.Vars(r)["zipcode"]

	if cep == "" || len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	resp, err := usecase.GetTemperatureByZipcode(cep)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
