package main

import (
	"github.com/corrots/oAuth/web"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", web.Index)

	r.Run(":8088")
}
