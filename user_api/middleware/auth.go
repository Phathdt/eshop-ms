package middleware

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "secret"

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

func GenerateJWT(id uint32) (string, error) {
	payload := ClaimsWithScope{}
	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
}
