package main

import (
	"os"

	"discord.land/site/server/utilities/oauth"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

var (
	clientID    = os.Getenv("CLIENT_ID")
	clientToken = os.Getenv("CLIENT_SECRET")
	base        = os.Getenv("BASE")
	port        = os.Getenv("PORT")
)

var (
	db *pg.DB
	oa *oauth.OAuth
)

func cors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Next()
}

func main() {
	oa = oauth.New(clientID, clientToken)
	println(oauth2Redirect)
	// var err error
	// db, err = database.Init()
	// if err != nil {
	// 	panic(err)
	// }

	app := gin.New()
	app.StaticFile("/", "../client/index.html")

	api := app.Group("/api")

	api.GET("/login", loginRoute)
	api.GET("/login/callback", loginCallbackRoute)

	app.Run(port)
}
