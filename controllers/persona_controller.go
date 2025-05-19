package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/danysoftdev/p-go-list/services"

)

func ObtenerPersonas(w http.ResponseWriter, r *http.Request) {
	personas, err := services.ListarPersonas()
	if err != nil {
		http.Error(w, "Error al obtener personas", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(personas)
}
