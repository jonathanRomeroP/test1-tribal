package routes

import (
	controller "test1-tribal/controllers"
	"test1-tribal/middleware"

	"github.com/gin-gonic/gin"
)

func SongRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate())
	//routes.GET("/song", controller.GetAllSong())
	routes.GET("/song/filter/", controller.FilterSong())
}
