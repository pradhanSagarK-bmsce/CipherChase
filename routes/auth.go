package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pradhanSagarK-bmsce/CipherChase/controllers"
)

func AdminRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/signup", controllers.SignUp)
	auth.POST("/signin", controllers.SignIn)

}
