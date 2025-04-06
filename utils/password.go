package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(c *gin.Context, password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return ""
	}
	return string(hashed)
}

func ComparePassword(hashedPassword string, password string) bool {

	matched := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return matched == nil
}
