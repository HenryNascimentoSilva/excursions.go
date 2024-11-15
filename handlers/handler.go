package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateExcursionHandler(ctx *gin.Context) {
 ctx.JSON(http.StatusOK, gin.H {
    "msg": "Excursion created",    
  }) 
}

func ShowExcursionHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H {
    "msg": "Showing Excursion",
  })
}

func DeleteExcursionHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H {
    "msg": "Delete excursion",
  })
}

func UpdateExcursionHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H {
    "msg": "Updating Excursion",
  })
}

func ShowExcursionsHandler(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H {
    "msg": "Showing all excursions",
  })
}
