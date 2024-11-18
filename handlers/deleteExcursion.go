package handler

import (
	"fmt"
	"net/http"

	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteExcursionHandler(ctx *gin.Context) {
  id := ctx.Query("id")
  if id == "" {
    sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
    return
  }
  
  excursion := schemas.Excursion{}

  if err := db.First(&excursion, id).Error; err != nil {
    sendError(ctx, http.StatusNotFound, fmt.Sprintf("excursion with id: %s", id))
    return
  }

  if err := db.Delete(&excursion).Error; err != nil {
    sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("erro deleting excursion with id: %s", id))
    return
  }


  ctx.JSON(http.StatusOK, gin.H {
    "msg": "excursion deleted successfully",
    "data": excursion,
  })
}