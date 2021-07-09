package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/models"
	"main.go/utils"
)

// r.POST("/auth/register", func(c *gin.Context) {
func Register(c *gin.Context) {
	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	useraddress := c.PostForm("useraddress")
	// password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	// _, err := database.DB.Query("Insert into users values ($1 ,$2, $3, )", email, password, phone)
	// if err != nil {
	// 	log.Fatal("Query error: ", err)
	// }
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
		Password:    password,
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

// r.GET("/get", func(c *gin.Context) {
// 	name, err := database.DB.Query("SELECT name FROM test WHERE id = 1")
// 	if err != nil {
// 		log.Fatal("Query error: ", err)
// 	}
// 	c.JSON(200, gin.H{
// 		"message": "name",
// 	})
// })

func Login(c *gin.Context) {
	// r.POST("/auth/login", func(c *gin.Context) {

	var user models.User
	email := c.PostForm("email")
	password := c.PostForm("password")
	database.DB.Where("email = ?", email).Find(&user)
	// fmt.Print(user.Email, user.Password, password)
	if user.Email == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	if user.Password == password {
		token, _ := utils.CreateJWT(email)
		c.JSON(200, gin.H{
			"message": token,
		})
		return
	}
	if user.Password != password {

		c.JSON(404, gin.H{
			"message": "Wrong password",
		})
		return
	}
}

func MailVerify(c *gin.Context) {
	// r.POST("/auth/verifyEmail", func(c *gin.Context) {

	var user models.User
	email := c.PostForm("email")
	database.DB.Where("email = ?", email).Find(&user)
	if user.Email == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}
	utils.Send(user.Email, "Verify Your Email Address", "Click here to verify:\n http://localhost:8080/auth/verify/"+email)

}
func VerifyEmail(c *gin.Context) {
	// r.GET("/auth/verify/:email", func(c *gin.Context) {

	var user models.User
	email := c.Param("email")
	database.DB.Where("email = ?", email).Find(&user)
	if user.Email == "" {
		c.JSON(404, gin.H{
			"message": "User not found",
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
	// r.GET("/auth/profile/:email", func(c *gin.Context) {

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

	token := c.PostForm("token")
	email := utils.ParseJWTToken(token)
	// var user models.User

	// database.DB.Where("email = ?", email).First(&user)

	c.JSON(200, gin.H{
		"message": email,
	})

}
