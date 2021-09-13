package autenticacao

import (
	"digital-bank/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CreateToken cria o token que será usado na autenticação
func CreateToken(accountID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["accountID"] = accountID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey))
}

//ValidarToken valida se o token da requisição é válido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)

	token, err := jwt.Parse(tokenString, retornarChaveVerificacao)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

//ExtrairAccountID extrai o token enviado na requisição
func ExtrairAccountID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveVerificacao)
	if err != nil {
		return 0, err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountID, err := strconv.ParseUint(fmt.Sprintf("%.f", permissoes["accountID"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return accountID, nil
	}

	return 0, errors.New("Token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornarChaveVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinaruta inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
