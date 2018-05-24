package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/bwmarrin/discordgo"

	"github.com/gin-gonic/gin"
)

var oauth2Redirect = discordgo.EndpointDiscord + "oauth2/authorize?client_id=" + clientID + "&scope=identify&response_type=code&callback_uri=" + url.QueryEscape(os.Getenv("BASE")+"/api/login/callback")

func loginRoute(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, oauth2Redirect)
}

func loginCallbackRoute(ctx *gin.Context) {
	code, exists := ctx.GetQuery("code")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "No code provided."})
		return
	}

	tr, err := oa.AuthorizeCode(code)
	if err != nil {
		log.Printf("Error while authorizing code: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to authorize given code."})
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
