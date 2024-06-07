package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pasin-mosps/ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engin) {
	incomingRoutes.POST("/users/singup", controllers.SingUp())
	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/user/search", controllers.SearchProductByQuery())
}
