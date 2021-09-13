package main

import (
	"digital-bank/src/config"
	"digital-bank/src/router"
	"digital-bank/src/utils"
	"fmt"
	"log"
	"net/http"
)

func init() {
	utils.CarregarTemplates()
}

func main() {
	config.CarregarVariaveisAmbiente()

	r := router.GerarRotas()

	fmt.Printf("Escutando na rodando na porta %d", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
