package controllers

import (
	"digital-bank/src/autenticacao"
	"digital-bank/src/db"
	"digital-bank/src/models"
	"digital-bank/src/repository"
	"digital-bank/src/resposta"
	"digital-bank/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAccounts retorna uma lista de contas
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	db, err := db.ConectDB()
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.AccountsRepository(db)

	accounts, err := repository.GetAccouts()
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resposta.JSON(w, http.StatusAccepted, accounts)

}

//GetBalance retorna o saldo de um conta
func GetBalance(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	accountID, err := strconv.ParseUint(parametros["account_id"], 10, 64)
	if err != nil {
		resposta.Erro(w, http.StatusBadRequest, err)
		return
	}

	accountIDToken, err := autenticacao.ExtrairAccountID(r)
	if err != nil {
		resposta.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if accountID != accountIDToken {
		resposta.Erro(w, http.StatusForbidden, errors.New("não é permitido obter o saldo de uma que não seja sua"))
		return
	}

	db, err := db.ConectDB()
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.AccountsRepository(db)

	balance, err := repository.GetAccountBalance(accountID)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resposta.JSON(w, http.StatusAccepted, balance)
}

//CreateAccount cria uma nova conta no banco de dados
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var account models.Account
	if err = json.Unmarshal(body, &account); err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	db, err := db.ConectDB()
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.AccountsRepository(db)

	secret, err := seguranca.Hash(account.Secret)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	account.Secret = string(secret)

	accountID, err := repository.CreateAccount(account)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	account.ID = accountID

	resposta.JSON(w, http.StatusCreated, account)
}

//DepositAccount faz o depósito em uma conta expecífica
func DepositAccount(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	deposit_amount, err := strconv.ParseFloat(parametros["deposit_amount"], 64)
	if err != nil {
		resposta.Erro(w, http.StatusBadRequest, err)
		return
	}

	accountIDToken, err := autenticacao.ExtrairAccountID(r)
	if err != nil {
		resposta.Erro(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.ConectDB()
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.AccountsRepository(db)

	err = repository.MakeDeposit(accountIDToken, deposit_amount)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resposta.JSON(w, http.StatusCreated, "Depósito efetuado com sucesso!")
}
