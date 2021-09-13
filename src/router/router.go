package router

import (
	"digital-bank/src/router/rotas"

	"github.com/gorilla/mux"
)

//GerarRotas retorna um router com as rotas da API
func GerarRotas() *mux.Router {
	r := mux.NewRouter()
	return rotas.ConfigRotas(r)
}
