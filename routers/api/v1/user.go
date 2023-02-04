package v1

import (
	"github.com/gin-gonic/gin"
)

func AddUserRoute(rg *gin.RouterGroup) {
	user := rg.Group("v1/user")

	user.POST("/", CreateUser)
}

func CreateUser(ctx *gin.Context) {
	// TODO: use go-playground/validator to validate
	var body struct {
		Email    string
		Password string
	}

	ctx.Bind(&body)
}
