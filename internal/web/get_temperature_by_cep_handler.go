package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetTemperatureByCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := mux.Vars(r)["cep"]

	w.Write([]byte(cep))
}
