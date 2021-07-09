package utils

import (
	"fmt"
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

func ParseJWTToken(tokenString string) string {
	// claim := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
		fmt.Print(err)
	}

	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer
}
