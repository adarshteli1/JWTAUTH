package services

import (
	"JWTAUTH/helpers"
	"JWTAUTH/models"
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func Signup(user models.User) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := validate.Struct(user); err != nil {
		return nil, err
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Panic(err)
		return nil, errors.New("error while checking for the email")
	}

	password := helpers.HashPassword(*user.Password)
	user.Password = &password

	count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	if err != nil {
		log.Panic(err)
		return nil, errors.New("error while checking for the phone")
	}

	if count > 0 {
		return nil, errors.New("this email or phone number already exists")
	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()

	token, refreshToken, _ := helpers.GenerateAllTokens(
		user.Email,
		*user.First_name,
		*user.Last_name,
		user.User_id,
		*user.User_type,
	)

	user.Token = &token
	user.Refresh_token = &refreshToken

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("User was not created")
	}

	return result, nil
}

func Login(user models.User) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var foundUser models.User

	err := userCollection.FindOne(ctx, bson.M{
		"email": user.Email,
	}).Decode(&foundUser)

	if err != nil {
		return models.User{}, errors.New("Email or Password is Incorrect")
	}

	passwordIsValid, msg := helpers.VerifyPassword(*foundUser.Password, *user.Password)

	if !passwordIsValid {
		return models.User{}, errors.New(msg)
	}

	if foundUser.Email == "" {
		return models.User{}, errors.New("User not found")
	}

	token, refreshToken, err := helpers.GenerateAllTokens(
		foundUser.Email,
		*foundUser.First_name,
		*foundUser.Last_name,
		foundUser.User_id,
		*foundUser.User_type,
	)

	if err != nil {
		return models.User{}, err
	}

	helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

	err = userCollection.FindOne(ctx, bson.M{
		"user_id": foundUser.User_id,
	}).Decode(&foundUser)

	if err != nil {
		return models.User{}, err
	}

	return foundUser, nil
}
