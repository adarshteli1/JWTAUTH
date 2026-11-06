package routes

import (
	"JWTAUTH/controllers"
	"JWTAUTH/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine,SecretKey string) {
	incomingRoutes.Use(middleware.Authenticate(SecretKey))
	incomingRoutes.GET("/users", middleware.AdminOnly(), controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", middleware.MatchUserID(), controllers.GetUser())

}
