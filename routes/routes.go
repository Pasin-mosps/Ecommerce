package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pasin-mosps/ecommerce/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users/signup", controllers.SignUp)
	router.POST("/user/login", controllers.Login)
	router.POST("/admin/addproduct", controllers.ProductViewAdmin)
	router.GET("/users/productview", controllers.SearchProduct)
	router.GET("/user/search", controllers.SearchProductByQuery)
}
