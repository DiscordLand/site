package main

import (
	"log"
	"net/http"

	"discord.land/site/server/database"
	"github.com/gin-gonic/gin"
)

func botGET(ctx *gin.Context) {
	var bot database.Bot
	err := db.Model(&bot).Where("bot.id = ?", ctx.Param("id")).Select()
	if err != nil {
		log.Println(err)
		sendMessage(ctx, http.StatusInternalServerError, "Failed to query database.")
		return
	}

	ctx.JSON(http.StatusOK, bot)
}

func botsGET(ctx *gin.Context) {
	var bots []database.Bot
	err := db.Model(&bots).Select()
	if err != nil {
		log.Println(err)
		sendMessage(ctx, http.StatusInternalServerError, "Failed to query database.")
		return
	}

	ctx.JSON(http.StatusOK, bots)
}

func botsPOST(ctx *gin.Context) {
	var bot database.Bot
	err := ctx.BindJSON(&bot)
	if err != nil {
		log.Println(err)
		sendMessage(ctx, http.StatusUnprocessableEntity, "Could not parse JSON.")
		return
	}

	if bot.Developers == nil {
		bot.Developers = []string{}
	}

	err = db.Insert(&bot)
	if err != nil {
		log.Println(err)
		sendMessage(ctx, http.StatusInternalServerError, "Failed to query database.")
		return
	}

	ctx.JSON(http.StatusOK, bot)
}
