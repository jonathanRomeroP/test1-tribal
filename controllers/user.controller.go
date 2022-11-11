package controllers

import (
	"context"
	"log"
	"net/http"
	"test1-tribal/database"
	"test1-tribal/helpers"
	"test1-tribal/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, foundUserPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(foundUserPassword), []byte(userPassword))
	check := true

	if err != nil {
		check = false
	}

	return check
}

func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario existente"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password
		user.ID = primitive.NewObjectID()

		token, refreshToken := helpers.GenerateAllTokens(*user.Email, *user.Name)
		user.Token = &token
		user.Refresh_token = &refreshToken

		resInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuario no creado"})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resInsertionNumber)
	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		passwordIsValid := VerifyPassword(*user.Password, *foundUser.Password)

		if !passwordIsValid || foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Datos incorrectos"})
			return
		}

		token, refreshToken := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.Name)
		helpers.UpdateAllTokens(token, refreshToken, *foundUser.Email)

		err = userCollection.FindOne(ctx, bson.M{"email": foundUser.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, foundUser)
	}
}
