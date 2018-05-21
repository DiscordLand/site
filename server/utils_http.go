package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendBadMessage(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": message})
}
