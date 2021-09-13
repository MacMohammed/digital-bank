package resposta

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON retorna as repostas no formato json
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

//Erro padroniza as resposta de erro que venham a acontecer. Retornando um json.
func Erro(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}
