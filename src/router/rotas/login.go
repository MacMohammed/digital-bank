package rotas

import (
	"digital-bank/src/controllers"
	"net/http"
)

var Login = Rota{

	URI:          "/login",
	Metodo:       http.MethodPost,
	Funcao:       controllers.Login,
	Autenticacao: false,
}
