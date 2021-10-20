package routes

import (
	controller "api_jwt_gin/controllers"
	"api_jwt_gin/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.DELETE("/users/:user_id", controller.DeleteUser())
	incomingRoutes.PUT("users/:user_id", controller.UpdateUser())
}
