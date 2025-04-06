package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/pradhanSagarK-bmsce/CipherChase/database"
	"github.com/pradhanSagarK-bmsce/CipherChase/models"
	"github.com/pradhanSagarK-bmsce/CipherChase/utils"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User

	if accessErr := c.BindJSON(&user); accessErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if utils.HashedPassword(c, user.Password) != "" {
		user.Password = utils.HashedPassword(c, user.Password)
	} else {
		return
	}

	// Insert into DB
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func SignIn(c *gin.Context) {
	var user models.User

	if accessErr := c.BindJSON(&user); accessErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userFromDB := models.User{}
	log.Println(user.Username)
	err := collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&userFromDB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}

	if !utils.ComparePassword(userFromDB.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Logged in successfully!",
		"access_token": token,
	})
}
