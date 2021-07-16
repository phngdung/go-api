package utils

import (
	"os"
	"strings"
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

func ParseJWTToken(bearerToken string) string {

	tokenString := strings.Split(bearerToken, " ")[1]

	var email string
	type MyCustomClaims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}

	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		email = claims.Email
	}
	return email

}
