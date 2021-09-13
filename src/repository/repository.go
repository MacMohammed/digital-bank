package repository

import "database/sql"

//É Repository é um tipo que armazena a conexão com o banco
type Repository struct {
	db *sql.DB
}
