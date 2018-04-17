package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.Writer.Write([]byte("ok"))
	})
	router.Run(":8080")
}
