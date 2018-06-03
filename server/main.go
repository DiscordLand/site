package main

import (
	"net/http"
	"os"

	"discord.land/site/server/database"
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
	var err error
	db, err = database.Init()
	if err != nil {
		panic(err)
	}

	app := gin.New()
	app.LoadHTMLGlob("site/server/client/*")
	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{})
	})
	// app.StaticFile("/", "../client/index.html")

	api := app.Group("/api")

	api.GET("/bots", botsGET)
	api.GET("/bots/:id", botGET)
	api.POST("/bots", botsPOST)
	// api.PATCH("/bots/:id", botsModifyRoute)
	// api.DELETE("/bots/:id", botsDeleteRoute)

	api.GET("/login", loginGET)
	api.GET("/login/callback", loginCallbackGET)
	api.POST("/logout", logoutPOST)

	app.Run(port)
}

type messageResponse struct {
	Message string `json:"message"`
}

func sendMessage(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, messageResponse{message})
}
