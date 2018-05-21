package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	api := app.Group("/api")
	api.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	})

	// users := api.Group("/users")
	// bots := api.Group("/bots")
	// users.GET("/")
	// bots.GET("/")

	api.GET("/login", loginRoute)
	api.GET("/login/callback", loginCallbackRoute)

	app.Run(port)
}
