package utils

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/pradhanSagarK-bmsce/CipherChase/models"
)

func GenerateToken(user models.User) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	mySigningKey := string(os.Getenv("JWT_SECRET"))
	// log.Println(mySigningKey)
	claims := &jwt.RegisteredClaims{
		Issuer:  "AppAdmin", // Token issuer
		Subject: user.ID.String(),
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// log.Println(token)
	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
