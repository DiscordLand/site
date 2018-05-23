package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

var db *pg.DB

func main() {
	// var err error
	// db, err = database.Init()
	// if err != nil {
	// 	panic(err)
	// }

	app := gin.New()
	app.StaticFile("/", "../client/index.html")

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
