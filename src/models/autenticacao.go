package models

type DadosAutenticacao struct {
	Name      string `json:"name"`
	AccountID string `json:"id"`
	Token     string `json:"token"`
}
