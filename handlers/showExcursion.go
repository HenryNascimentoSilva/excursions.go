package handler

import (
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

func ShowExcursionHandler(ctx *gin.Context) {
  id := ctx.Query("id")

  excursion := schemas.Excursion{}

  if err := db.First(&excursion, id).Error; err != nil {
    sendError(ctx, http.StatusNotFound, "excursion not found!")
    return
  }

  ctx.JSON(http.StatusOK, gin.H {
    "data": excursion,
  })
}