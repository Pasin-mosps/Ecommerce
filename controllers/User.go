package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pasin-mosps/ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	message := ""
	if err != nil {
		message = "username or password is Incorrect"
		valid = false
	}
	return valid, message
}

func SignUp(c *gin.Context) {
	context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(htp.StatusBadRequest, gin.H{"error": err})
		return
	}

	validationErr := Validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	if count, err := UserCollection.CountDocuments(context, bson.M{"email": user.Email}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.ERROR()})
		return
	} else if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword := HashPassword(*user.Password)
	user.Password = &hashedPassword

	currentTime := time.Now().Format(time.RFC3339)
	user.Created_At, _ = time.Parse(time.RFC3339, currentTime)
	user.Updated_At, _ = time.Parse(time.RFC3339, currentTime)
	user.ID = primitive.NewObjectID()
	user.User_ID = user.ID.Hex()
	token, refreshToken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
	user.Token = &token
	user.Refresh_Token = &refreshToken

	user.UserCart = []models.ProductUser{}
	user.Address_Details = []models.Address{}
	user.Order_Status = []models.Order{}

	if _, err := UserCollection.InsertOne(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not created"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully Signed Up!!"})
}

func Login(c *gin.Context) {
	var context, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User
	var foundUser models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := UserCollection.FindOne(context, bson.M{"email": user.Email}).Decode((&foundUser))
	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username or password incorrect"})
		return
	}

	passwordIsValid, message := VerifyPassword(*user.Password, *foundUser.Password)
	defer cancel()

	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		fmt.Println(message)
		return
	}

	token, refreshToken, _ := generate.TokenGenerator(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name, *foundUser.User_ID)
	defer cancel()
	generate.UpdateAllTokens(token, refreshToken, foundUser.User_ID)
	c.JSON(http.StatusFound, foundUser)
}

func LogOut(c *gin.Context) {
	return
}
