package helpers

import (
	"api_jwt_gin/database"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type SignedDetails struct {
	FirstName string
	LastName  string
	Email     string
	UserType  string
	UserID    string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY []byte

func GenerateAllTokens(email string, firstName string, lastName string, userType string, userID string) (signedToken string, signedRefreshToken string, err error) {
	err1 := godotenv.Load("mongo.env")
	if err1 != nil {
		log.Fatal("Error loading the mongo.env file.")
	}

	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
	claims := &SignedDetails{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		UserType:  userType,
		UserID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println(token1)
	token, err := token1.SignedString(SECRET_KEY)
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(SECRET_KEY)
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, foundUserID string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	upsert := true

	filter := bson.M{"user_id": foundUserID}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)

	defer cancel()

	if err != nil {
		log.Panicln(err)
	}
	return
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return SECRET_KEY, nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	fmt.Println(claims)
	if !ok {
		msg = fmt.Sprintf("The token entered is invalid")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("The token entered has expired.")
		msg = err.Error()
		return
	}
	fmt.Println(claims, msg)
	return claims, msg
}
