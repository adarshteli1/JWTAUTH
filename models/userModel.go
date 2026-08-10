package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Email         string             `json:"email" bson:"email" validate:"email,required"`
	Password      *string            `json:"password" bson:"password" validate:"required,min=6"`
	Phone         *string            `json:"phone" validate:"required"`
	User_type     *string            `json:"user_type" validate:"required,oneof=ADMIN USER"`
	User_id       string             `json:"user_id"`
	Refresh_token *string            `json:"-" bson:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}

type UserResponse struct {
	ID         primitive.ObjectID `json:"id"`
	First_name *string            `json:"first_name"`
	Last_name  *string            `json:"last_name" `
	Email      string             `json:"email"`
	Phone      *string            `json:"phone" `
	User_type  *string            `json:"user_type" `
	User_id    string             `json:"user_id"`
}
