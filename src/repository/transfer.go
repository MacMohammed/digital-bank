package repository

import (
	"context"
	"database/sql"
	"digital-bank/src/models"
	"errors"
)

func TransferRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

//CreateTransfer cria uma transfereência entre contas
func (repository Repository) CreateTransfer(transfer models.Transfer, account_origin_id uint64) (uint64, error) {
	ctx := context.Background()

	tx, err := repository.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var balance float64
	err = tx.QueryRowContext(ctx, "select balance from accounts where id = $1", account_origin_id).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if balance < transfer.Amount {
		tx.Rollback()
		return 0, errors.New("saldo insuficiente")
	}

	_, err = tx.ExecContext(ctx, "update accounts set balance = balance - $1 where id = $2", transfer.Amount, transfer.AccountOriginID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.ExecContext(ctx, "update accounts set balance = (balance + $1) where id = $2", transfer.Amount, transfer.AccountDestinationID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var transferID uint64
	err = tx.QueryRowContext(ctx, "insert into transfers (account_origin_id, account_destination_id, amount) values ($1, $2, $3) RETURNING id;", transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount).Scan(&transferID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint64(transferID), nil
}

//GetAccountTransfers retorna as transferências que um id fez
func (repository Repository) GetAccountTransfers(accountID uint64) ([]models.Transfer, error) {
	rows, err := repository.db.Query(`
		select 
			tr.id,
			tr.account_origin_id,
			tr.account_destination_id,
			ac.name,
			tr.amount,
			tr.created_at
		from transfers tr join accounts ac
			on ac.id = tr.account_destination_id
		where 
			tr.account_origin_id = $1
		order by 
			ac.name,
			tr.created_at`, accountID)

	if err != nil {
		return []models.Transfer{}, err
	}

	defer rows.Close()

	var transfers []models.Transfer

	for rows.Next() {
		var transfer models.Transfer

		if err := rows.Scan(
			&transfer.ID,
			&transfer.AccountOriginID,
			&transfer.AccountDestinationID,
			&transfer.Name,
			&transfer.Amount,
			&transfer.Created_at,
		); err != nil {
			return []models.Transfer{}, err
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}
