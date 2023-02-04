package main

import (
	docs "go-backend/docs"
	initialize "go-backend/initialize"
	routers "go-backend/routers"

	http "net/http"

	gin "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initialize.LoadEnv()
	initialize.ConnectToDB()
	initialize.MigrateDB()
}

// @title			Go Backend
// @version		0.1
// @description	This is a sample server implement with Go.
// @schemes		http,https
// @host			localhost:8080
// @contact.name	Eddie Lin
// @license.name	MIT
// @license.url	https://opensource.org/licenses/MIT
// @BasePath		/api/v1/
func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	// API Docs
	docs.SwaggerInfo.BasePath = "/api"
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.GET("/", func(ctx *gin.Context) { ctx.Redirect(http.StatusPermanentRedirect, "/docs/index.html") })

	routers.SetupRouter(app)
	app.Run()
}
