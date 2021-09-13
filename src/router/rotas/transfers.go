package rotas

import (
	"digital-bank/src/controllers"
	"net/http"
)

var Transfers = []Rota{
	{
		URI:          "/transfers",
		Metodo:       http.MethodGet,
		Funcao:       controllers.GetTransfers,
		Autenticacao: true,
	},
	{
		URI:          "/transfers",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CreateTransfer,
		Autenticacao: true,
	},
}
