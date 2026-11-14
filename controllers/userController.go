package controllers

import (
	services "JWTAUTH/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		user, err := services.GetUser(userId, c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
