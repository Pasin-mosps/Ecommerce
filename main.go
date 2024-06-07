package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/pasin-mosps/ecommerce/controllers"
	// "github.com/pasin-mosps/ecommerce/database"
	// "github.com/pasin-mosps/ecommerce/middleware"
	"github.com/pasin-mosps/ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserDData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	// router.Use(middileware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveIteem())
	router.GET("/cartcheckout", app.ButFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
