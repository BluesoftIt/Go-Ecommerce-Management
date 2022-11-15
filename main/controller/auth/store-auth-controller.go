package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreAuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type storeAuthController struct {
}

func NewStoreAuthController() StoreAuthController {
	return &storeAuthController{}
}

func (c *storeAuthController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
	})
}

func (c *storeAuthController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}
