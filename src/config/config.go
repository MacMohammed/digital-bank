package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConnectionDataBase é a string de conexão com o banco
	StringConnectionDataBase = ""

	//Driver representa o drive usado para conectar no banco de dados
	Driver = ""

	//ApiPort é a porta padrão da aplicação
	ApiPort = 0

	//SecretKey é a chave que será usada no token
	SecretKey []byte
)

//CarregarVariaveisAmbiente carrega as variáveis de ambiente da aplicação
func CarregarVariaveisAmbiente() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPort = 9000
	}

	StringConnectionDataBase = fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	Driver = os.Getenv("POSTGRES_DRIVER")

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
