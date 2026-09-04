package services

import (
	"JWTAUTH/database"
	"JWTAUTH/models"
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpeCollection(database.Client, "user")

func GetUsers(c *gin.Context) ([]bson.M, error) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 10
	}
	if recordPerPage > 100 {
		recordPerPage = 100
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordPerPage

	matchStage := bson.D{
		{Key: "$match", Value: bson.D{}},
	}
	projectUserStage := bson.D{
		{
			Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "id", Value: 1},
				{Key: "first_name", Value: 1},
				{Key: "last_name", Value: 1},
				{Key: "email", Value: 1},
				{Key: "phone", Value: 1},
				{Key: "user_type", Value: 1},
			},
		},
	}
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "total_count", Value: bson.D{
				{Key: "$sum", Value: 1},
			}},
			{Key: "data", Value: bson.D{
				{Key: "$push", Value: "$$ROOT"},
			}},
		}},
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: 1},
			{Key: "user_items", Value: bson.D{
				{Key: "$slice", Value: []interface{}{
					"$data",
					startIndex,
					recordPerPage,
				}},
			}},
		}},
	}

	result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		projectUserStage,
		groupStage,
		projectStage,
	})

	if err != nil {
		return nil, err
	}

	var allUsers []bson.M

	if err = result.All(ctx, &allUsers); err != nil {
		return nil, err
	}

	return allUsers, nil
}

func GetUser(userId string, c *gin.Context) (models.User, error) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 100*time.Second)
	defer cancel()

	var user models.User

	err := userCollection.FindOne(ctx, bson.M{
		"user_id": userId,
	}).Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
