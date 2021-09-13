package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash receber a senha string e convert em um hash de bytes.
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerificarSenha compara as strings da senha e o hash armazenado no banco de dados
func VerificarSenha(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
