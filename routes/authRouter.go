package routes

import (
	controller "JWTAUTH/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine,SecretKey string) {
	incomingRoutes.POST("users/signup", controller.Signup(SecretKey))
	incomingRoutes.POST("users/login", controller.Login(SecretKey))
}
