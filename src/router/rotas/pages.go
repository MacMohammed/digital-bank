package rotas

import (
	"digital-bank/src/pages"
	"net/http"
)

var Pages = []Rota{
	{
		URI:          "/",
		Metodo:       http.MethodGet,
		Funcao:       pages.CarregarTelaCadastro,
		Autenticacao: false,
	},
	{
		URI:          "/login",
		Metodo:       http.MethodGet,
		Funcao:       pages.CarregarTelaLogin,
		Autenticacao: false,
	},
}
