package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"discord.land/site/server/database"

	"github.com/bwmarrin/discordgo"

	"github.com/gin-gonic/gin"
)

var oauth2Redirect = discordgo.EndpointDiscord + "oauth2/authorize?client_id=" + clientID + "&scope=identify&response_type=code&callback_uri=" + url.QueryEscape(os.Getenv("BASE")+"/api/login/callback")

func loginGET(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, oauth2Redirect)
}

func loginCallbackGET(ctx *gin.Context) {
	code, exists := ctx.GetQuery("code")
	if !exists {
		sendMessage(ctx, http.StatusBadRequest, "No code provided.")
		return
	}

	tr, err := oa.AuthorizeCode(code)
	if err != nil {
		log.Printf("Error while authorizing code: %v\n", err)
		sendMessage(ctx, http.StatusBadRequest, "Failed to authorize given code.")
		return
	}

	// du, err := oa.FetchUser(tr.AccessToken)
	// if err != nil {
	// 	panic(err)
	// }

	t, _ := json.Marshal(tr)

	ctx.HTML(http.StatusOK, "callback", gin.H{"tr": string(t)})
}

func logoutPOST(ctx *gin.Context) {
	data := database.Session{}
	err := ctx.BindJSON(&data)
	if err != nil {
		log.Println(err)
		sendMessage(ctx, http.StatusUnprocessableEntity, "Could not parse JSON.")
		return
	}

	if data.Token == "" {
		sendMessage(ctx, http.StatusUnprocessableEntity, "Token not proivded.")
		return
	}

	db.Delete(&data)
}
