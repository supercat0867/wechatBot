package main

import (
	"github.com/gin-gonic/gin"
	"wechatBot/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.RegisterMessageHandler(r)

	r.Run(":8080")
}
