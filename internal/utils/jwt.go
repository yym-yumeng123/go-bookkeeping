package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

var hmacSampleSecret = []byte("yym1994010203wxf1234567890")

// 生成 jwt
func GenerateJWT(id int) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
	})

	return token.SignedString(hmacSampleSecret)
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

func Parse(jwtString string) (*jwt.Token, error) {
	return jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
}
