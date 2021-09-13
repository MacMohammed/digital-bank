package rotas

import (
	"digital-bank/src/controllers"
	"net/http"
)

var Accounts = []Rota{
	{
		URI:          "/accounts",
		Metodo:       http.MethodGet,
		Funcao:       controllers.GetAccounts,
		Autenticacao: false,
	},
	{
		URI:          "/accounts/{account_id}/balance",
		Metodo:       http.MethodGet,
		Funcao:       controllers.GetBalance,
		Autenticacao: true,
	},
	{
		URI:          "/accounts",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CreateAccount,
		Autenticacao: false,
	},
	{
		URI:          "/accounts/{deposit_amount}/deposit",
		Metodo:       http.MethodPost,
		Funcao:       controllers.DepositAccount,
		Autenticacao: true,
	},
}
