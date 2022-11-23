package main

import (
	"context"
	"net/http"
	"test1-tribal/config"
	"test1-tribal/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ctx         context.Context
	mongoclient *mongo.Client
)

func init() {

	/* 	ctx = context.TODO()

	   	// Connect to MongoDB
	   	mongoconn := options.Client().ApplyURI(Config.DBUri)
	   	mongoclient, err := mongo.Connect(ctx, mongoconn)

	   	if err != nil {
	   		panic(err)
	   	}

	   	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
	   		panic(err)
	   	}

	   	fmt.Println("MongoDB successfully connected...") */
}

func main() {
	var config, _ = config.LoadConfig(".")
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Bienvenido v2"})
	})
	routes.AuthRoutes(router)
	routes.SongRoutes(router)
	router.Run(":" + config.Port)
}
