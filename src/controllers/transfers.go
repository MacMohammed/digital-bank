package controllers

import (
	"digital-bank/src/autenticacao"
	"digital-bank/src/db"
	"digital-bank/src/models"
	"digital-bank/src/repository"
	"digital-bank/src/resposta"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//GetTransfers retorna a lista de transferencias do usuario autenticada.
func GetTransfers(w http.ResponseWriter, r *http.Request) {
	db, err := db.ConectDB()
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	accountID, err := autenticacao.ExtrairAccountID(r)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	transferRepository := repository.TransferRepository(db)

	transfers, err := transferRepository.GetAccountTransfers(accountID)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	resposta.JSON(w, http.StatusAccepted, transfers)
}

//Função resposável por receber a requisição para a criação de uma nova conta
func CreateTransfer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var transfer models.Transfer
	if err = json.Unmarshal(body, &transfer); err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	//Extração do id da conta do usuário logado na aplicação
	account_origin_id, err := autenticacao.ExtrairAccountID(r)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	transfer.AccountOriginID = account_origin_id

	db, err := db.ConectDB()
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	transferRepository := repository.TransferRepository(db)
	transferID, err := transferRepository.CreateTransfer(transfer, account_origin_id)
	if err != nil {
		resposta.Erro(w, http.StatusInternalServerError, err)
		return
	}

	transfer.ID = transferID
	resposta.JSON(w, http.StatusCreated, transfer)
}
