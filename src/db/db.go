package db

import (
	"database/sql"
	"digital-bank/src/config"

	_ "github.com/lib/pq"
)

//ConectDB devolve a conexão com o banco de dados, ou um erro.
func ConectDB() (*sql.DB, error) {
	db, err := sql.Open(config.Driver, config.StringConnectionDataBase)
	if err != nil {
		return nil, err
	}

	//Se ocorrer qualquer erro ao tentar pingar o banco,
	//fehcha-se a conecção e retorna o erro.
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	//Caso contrário, retorna a conexão aberta e um vazio no lugar do erro
	return db, nil
}
