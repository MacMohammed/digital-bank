package pages

import (
	"digital-bank/src/utils"
	"net/http"
)

func CarregarTelaCadastro(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "index.html", nil)
}

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}
