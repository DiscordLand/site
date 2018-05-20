package main

import "github.com/gin-gonic/gin"

const port = ":8002"

func main() {
	app := gin.New()

	app.StaticFile("/", "./index.html")

	app.Run(port)
}
