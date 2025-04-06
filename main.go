package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pradhanSagarK-bmsce/CipherChase/database"
	"github.com/pradhanSagarK-bmsce/CipherChase/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	fmt.Println("Port : ", port)

	r := gin.Default()
	database.ConnectDB()
	routes.AdminRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running!"})
	})

	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run(":8080")
	}
}
