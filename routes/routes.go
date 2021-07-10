package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func Setup(r *gin.Engine) {

	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/verifyEmail", controllers.VerifyEmail)
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/profile", controllers.Profile)
}
