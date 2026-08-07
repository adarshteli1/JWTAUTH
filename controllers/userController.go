package controllers

import (
	"JWTAUTH/database"
	"JWTAUTH/helpers"
	services "JWTAUTH/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpeCollection(database.Client, "user")
var validate = validator.New()

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := helpers.CheckUserType(c, "ADMIN"); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		users, err := services.GetUsers(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		user, err := services.GetUser(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
