package routers

import (
	v1 "go-backend/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	apiRoute := app.Group("api")
	v1.AddUserRoute(apiRoute)
}
