package models

import "time"

type Account struct {
	ID         uint64    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CPF        string    `json:"cpf,omitempty"`
	Secret     string    `json:"secret,omitempty"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at,omitempty"`
}
