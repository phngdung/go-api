package controllers

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/database"
	"main.go/models"
	"main.go/utils"
)

func Register(c *gin.Context) {
	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	useraddress := c.PostForm("useraddress")
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	database.DB.Where("email = ?", email).Find(&user)
	fmt.Print(user.Email)
	if user.Email != "" {
		c.JSON(404, gin.H{
			"message": "Account with email : " + email + " is already exist",
		})
		return
	}
	user = models.User{
		Email:       email,
		Password:    string(pw),
		Phone:       phone,
		Useraddress: useraddress,
		Status:      "notVerfied",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	database.DB.Create(&user)
	c.JSON(200, gin.H{
		"message": user,
	})

}

func Login(c *gin.Context) {

	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")
	database.DB.Where("email = ?", email).Find(&user)
	check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if check != nil {
		c.JSON(404, gin.H{
			"message": "Email or Password incorrect",
		})
		return
	}
	token, err := utils.CreateJWT(email)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": token,
	})

}

func VerifyEmail(c *gin.Context) {

	var user models.User
	email := c.PostForm("email")
	database.DB.Where("email = ?", email).Find(&user)
	if user.Email == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	if err := database.DB.Model(&user).Where("email = ?", email).Updates(models.User{VerifyCode: utils.RandStringBytes(8), VerifyExp: time.Now().Add(time.Hour * 3)}); err != nil {
		fmt.Print(err)
		c.JSON(404, gin.H{
			"message": "Error",
		})
	}

	utils.Send(user.Email, "Verify Your Email Address", "Click here to verify:\n http://localhost:8080/auth/verify/"+email+"/code/"+user.VerifyCode)

}
func Verify(c *gin.Context) {

	var user models.User
	email := c.Param("email")
	code := c.Param("verifyCode")
	database.DB.Where("email = ?", email).Find(&user)
	if user.Email == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}
	if user.Status == "verified" {
		c.JSON(404, gin.H{
			"message": "User verified",
		})
		return
	}
	if user.VerifyCode != code || user.VerifyExp.Before(time.Now()) {
		fmt.Print(user.VerifyCode, code, user.VerifyExp.Before(time.Now()))
		c.JSON(404, gin.H{
			"message": "Cannot Verify",
		})
		return
	}
	if err := database.DB.Model(&user).Where("email = ?", email).Update("status", "verified").Error; err != nil {
		fmt.Print(err)
		c.JSON(404, gin.H{
			"message": "Error" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
func GetProfile(c *gin.Context) {

	var user models.User
	email := c.Param("email")
	database.DB.Where("email = ?", email).Find(&user)
	if user.Email == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": user,
	})

}
func Profile(c *gin.Context) {
	// tokenString, _ := c.Cookie("ACCESS-KEY")

	tokenString := c.PostForm("token")
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
	var user models.User

	database.DB.Where("email = ?", email).First(&user)
	if user.Email == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": user,
	})

}
