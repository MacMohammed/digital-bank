package controllers

import (
	"digital-bank/src/autenticacao"
	"digital-bank/src/db"
	"digital-bank/src/models"
	"digital-bank/src/repository"
	"digital-bank/src/resposta"
	"digital-bank/src/seguranca"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Login é função responsável por receber a requiseição de login e encaminhar para o repositório de login.
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var account models.Account
	if err := json.Unmarshal(body, &account); err != nil {
		resposta.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectDB()
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.AccountsRepository(db)
	accountByCPF, err := repository.GetAccountByCPF(account.CPF)

	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := seguranca.VerificarSenha(accountByCPF.Secret, account.Secret); err != nil {
		resposta.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := autenticacao.CreateToken(accountByCPF.ID)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	var dadosAutenticacao models.DadosAutenticacao
	dadosAutenticacao.Name = accountByCPF.Name
	dadosAutenticacao.AccountID = fmt.Sprintf("%d", accountByCPF.ID)
	dadosAutenticacao.Token = token

	resposta.JSON(w, http.StatusOK, dadosAutenticacao)
}
