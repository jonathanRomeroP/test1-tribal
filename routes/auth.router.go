package routes

import (
	controller "test1-tribal/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("user/signup", controller.Signup())
	routes.POST("user/login", controller.Login())
}
