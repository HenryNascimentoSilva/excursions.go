package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
  // initializing router
  r := gin.Default()
  // defining a route
  r.GET("/ping", func(ctx *gin.Context) {
  // return context
    ctx.JSON(200, gin.H{
      "message:": "pong",
    })
  })
  r.Run()
}
