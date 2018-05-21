package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginRoute(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, discordTokenRedirect)
}

func loginCallbackRoute(ctx *gin.Context) {
	code, exists := ctx.GetQuery("code")
	if !exists {
		sendBadMessage(ctx, "No code parameter was provided.")
		return
	}

	tr, err := getToken(code)
	if err != nil {
		sendBadMessage(ctx, "An error occurred while attempting to authorize provided code.")
		return
	}

	fmt.Printf("AccessToken: %s\nTokenType: %s\nExpiresIn: %v\nRefreshToken: %s\nScope: %s\n", tr.AccessToken, tr.TokenType, tr.ExpiresIn, tr.RefreshToken, tr.Scope)
}
