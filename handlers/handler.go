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
  id := ctx.Param("id")
  ctx.JSON(http.StatusOK, gin.H {
    "msg": "Delete excursion",
    "id": id,
  })
}

func UpdateExcursionHandler(ctx *gin.Context) {
  id := ctx.Param("id")
  ctx.JSON(http.StatusOK, gin.H {
    "msg": "Updating Excursion",
    "id": id,
  })
}

func ShowExcursionsHandler(ctx *gin.Context) {
  id := ctx.Param("id")
  ctx.JSON(http.StatusOK, gin.H {
    "msg": "Showing all excursions",
    "id": id,
  })
}
