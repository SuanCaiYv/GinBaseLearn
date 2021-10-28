package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test1", func(context *gin.Context) {
		context.Set("key", "msl")
		context.String(200, "ok")
	})
	r.GET("/test2", func(context *gin.Context) {
		fmt.Println(context.GetString("key"))
		context.String(200, "ok")
	})
	r.Run(":8190")
}
