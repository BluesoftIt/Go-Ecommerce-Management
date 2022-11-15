package main

import (
	"github.com/bluesoftit/go-api/config"
	"github.com/bluesoftit/go-api/main/controller/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                 = config.Connect()
	storeAuthController auth.StoreAuthController = auth.NewStoreAuthController()
)

func main() {
	defer config.DisConnect(db)
	r := gin.Default()

	// initialize group of authentication route
	authRoute := r.Group("api/auth")
	{
		authRoute.POST("/login", storeAuthController.Login)
		authRoute.POST("/register", storeAuthController.Register)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
