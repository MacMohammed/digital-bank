package rotas

import (
	"digital-bank/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa a estrutura das rotas da api
type Rota struct {
	URI          string
	Metodo       string
	Funcao       func(http.ResponseWriter, *http.Request)
	Autenticacao bool
}

//ConfigRotas coloca todas as rotas dentro do router
func ConfigRotas(router *mux.Router) *mux.Router {

	rotas := Accounts
	rotas = append(rotas, Login)
	rotas = append(rotas, Transfers...)
	rotas = append(rotas, Pages...)

	for _, rota := range rotas {
		if rota.Autenticacao {
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))

	return router
}
