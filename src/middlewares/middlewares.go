package middlewares

import (
	"digital-bank/src/autenticacao"
	"digital-bank/src/resposta"
	"net/http"
)

//Logger garante que as rotas que precisarem de autenticação seja devidamente verificadas.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

//Autenticar é a função entre a requisição e a função que valida o token da requisição
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			resposta.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
