package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func Setup(r *gin.Engine) {

	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/verifyEmail", controllers.VerifyEmail)
	r.POST("/auth/login", controllers.Login)
	r.GET("/auth/profile", controllers.Profile)
	r.GET("/auth/verify/:email/code/:verifyCode", controllers.Verify)
}
