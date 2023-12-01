package utils

import (
	"crypto/rand"

	"github.com/golang-jwt/jwt/v5"
)

// 生成 jwt
func GenerateJWT(user_id int) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
	})

	key, err := generateHMACKey()

	if err != nil {
		return "", err
	}

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(key)
}

// 生成 key
func generateHMACKey() ([]byte, error) {
	// 秘钥长度 64 字节
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil

}
