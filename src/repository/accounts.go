package repository

import (
	"context"
	"database/sql"
	"digital-bank/src/models"
)

func AccountsRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

//CreateAccount insere uma nova conta no banco de dados
func (repository Repository) CreateAccount(account models.Account) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO accounts (name, cpf, secret) values ($1, $2, $3) RETURNING id;")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	var accountID uint64

	err = statement.QueryRow(account.Name, account.CPF, account.Secret).Scan(&accountID)
	if err != nil {
		return 0, err
	}

	return uint64(accountID), nil
}

//GetAccountByCPF retorna os dados de uma conta filtrado pelo cpf
func (repository Repository) GetAccountByCPF(cpf string) (models.Account, error) {

	var accounts models.Account

	defer repository.db.Close()

	err := repository.db.QueryRow("select id, name, cpf, secret, balance, created_at from accounts where cpf = $1;", cpf).Scan(
		&accounts.ID,
		&accounts.Name,
		&accounts.CPF,
		&accounts.Secret,
		&accounts.Balance,
		&accounts.Created_at,
	)
	if err != nil {
		return models.Account{}, err
	}

	return accounts, nil
}

//GetAccountBalance retorna o saldo de uma conta de um id expecifico
func (repository Repository) GetAccountBalance(accountID uint64) (float64, error) {
	var balance float64

	err := repository.db.QueryRow("Select balance from accounts where id = $1", accountID).Scan(&balance)
	if err != nil {
		return 0, err
	}

	defer repository.db.Close()

	return float64(balance), nil
}

//GetAccouts retorna todas as contas cadastradas no banco de dados
func (repository Repository) GetAccouts() ([]models.Account, error) {
	rows, err := repository.db.Query(`
		select
			id,
			name,
			cpf,
			secret,
			balance,
			created_at
		from
			accounts;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accouts []models.Account

	for rows.Next() {
		var account models.Account

		if err := rows.Scan(
			&account.ID,
			&account.Name,
			&account.CPF,
			&account.Secret,
			&account.Balance,
			&account.Created_at,
		); err != nil {
			return nil, err
		}

		accouts = append(accouts, account)
	}

	return accouts, nil
}

//MakeDeposit faz um depósito em uma conta de id expecífico
func (repository Repository) MakeDeposit(accountID uint64, amount float64) error {
	ctx := context.Background()

	tx, err := repository.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "insert into deposits (account_origin_id, amount) values ($1, $2);", accountID, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "update accounts set balance = balance+$1 where id = $2;", amount, accountID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
