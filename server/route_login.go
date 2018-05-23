package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginRoute(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, discordOauth2Redirect)
}

func loginCallbackRoute(ctx *gin.Context) {
	code, exists := ctx.GetQuery("code")
	if !exists {
		sendBadMessage(ctx, "No code parameter was provided.")
		return
	}

	tr, err := authorizeOauth2Code(code)
	if err != nil {
		log.Printf("Error while authorizing code: %v\n", err)
		sendBadMessage(ctx, "An error occurred while attempting to authorize provided code.")
		return
	}

	// ctx.Header("Content-Type", "application/json")
	// ctx.Writer.Write(body)
	// du, err := fetchDiscordUser(tr.AccessToken)
	// if err != nil {
	// 	panic(err)
	// }

	ctx.JSON(http.StatusOK, tr)
}
