package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_Name    string             `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name     string             `json:"last_name" validate:"required,min=2,max=100"`
	Email         string             `json:"email" validate:"required,min=4"`
	Phone_Number  string             `json:"phone_number" validate:"required,min=6,max=10"`
	Password      string             `json:"password" validate:"required,min=6"`
	User_Type     string             `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Token         string             `json:"token"`
	Refresh_Token string             `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_ID       string             `json:"user_id"`
}
