package helpers

import (
	"context"
	"log"
	"test1-tribal/config"
	"test1-tribal/database"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email string
	Name  string
	jwt.StandardClaims
}

var configData, _ = config.LoadConfig("./")
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var SECRET_KEY string = configData.AccessTokenPrivateKey

func GenerateAllTokens(email string, name string) (signedToken string, signedRefreshToken string) {

	/* claims := &SignedDetails{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "test",
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(configData.AccessTokenMaxAge)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(configData.RefreshTokenMaxAge)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte("mysecretkey"))
	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))
	*/

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    email,
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(configData.AccessTokenMaxAge)).Unix(),
	})

	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    email,
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(configData.RefreshTokenMaxAge)).Unix(),
	})

	token, err := claims.SignedString([]byte(SECRET_KEY))
	refreshToken, refErr := refreshClaims.SignedString([]byte(SECRET_KEY))

	if err != nil || refErr != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken
}

func UpdateAllTokens(token string, refreshToken string, email string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", token})
	updateObj = append(updateObj, bson.E{"refresh_token", refreshToken})

	upsert := true
	filter := bson.M{"email": email}
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
		log.Panic(err)
	}
	return
}

func ValidateToken(clientToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		clientToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		msg = "Token invalido"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "El token caducado"
	}

	return
}
