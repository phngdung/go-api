package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateJWT(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() //Token hết hạn sau 3 giờ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
