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

type AuthResponse struct {
	User         models.UserResponse `json:"user"`
	Token        string              `json:"token"`
	RefreshToken string              `json:"refresh_token"`
}

var validate = validator.New()

func Signup(user models.User) (AuthResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := validate.Struct(user); err != nil {
		return AuthResponse{}, err
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Panic(err)
		return AuthResponse{}, errors.New("error while checking for the email")
	}

	password := helpers.HashPassword(*user.Password)
	user.Password = &password

	count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	if err != nil {
		log.Panic(err)
		return AuthResponse{}, errors.New("error while checking for the phone")
	}

	if count > 0 {
		return AuthResponse{}, errors.New("this email or phone number already exists")
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

	user.Refresh_token = &refreshToken

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return AuthResponse{}, errors.New("User was not created")
	}

	userResponse := models.UserResponse{
		ID:         user.ID,
		First_name: user.First_name,
		Last_name:  user.Last_name,
		Email:      user.Email,
		Phone:      user.Phone,
		User_type:  user.User_type,
		User_id:    user.User_id,
	}

	return AuthResponse{
		User:         userResponse,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func Login(user models.User) (AuthResponse, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var foundUser models.User

	err := userCollection.FindOne(ctx, bson.M{
		"email": user.Email,
	}).Decode(&foundUser)

	if err != nil {
		return AuthResponse{}, errors.New("Email or Password is Incorrect")
	}

	passwordIsValid, msg := helpers.VerifyPassword(*foundUser.Password, *user.Password)

	if !passwordIsValid {
		return AuthResponse{}, errors.New(msg)
	}

	if foundUser.Email == "" {
		return AuthResponse{}, errors.New("User not found")
	}

	token, refreshToken, err := helpers.GenerateAllTokens(
		foundUser.Email,
		*foundUser.First_name,
		*foundUser.Last_name,
		foundUser.User_id,
		*foundUser.User_type,
	)

	if err != nil {
		return AuthResponse{}, err
	}

	helpers.UpdateAllTokens(refreshToken, foundUser.User_id)

	userResponse := models.UserResponse{
		ID:         foundUser.ID,
		First_name: foundUser.First_name,
		Last_name:  foundUser.Last_name,
		Email:      foundUser.Email,
		Phone:      foundUser.Phone,
		User_type:  foundUser.User_type,
		User_id:    foundUser.User_id,
	}

	return AuthResponse{
		User:         userResponse,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
