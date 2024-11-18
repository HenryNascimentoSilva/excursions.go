package handler

import (
	"net/http"

	"excursion.com/config"
	"excursion.com/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
  logger *config.Logger
  db *gorm.DB
)




func UpdateExcursionHandler(ctx *gin.Context) {
  req := UpdateExcursionRequest{}

  ctx.BindJSON(&req)

  if err := req.Validate(); err != nil {
    logger.Errorf("validation erro: %v", err.Error())
    sendError(ctx, http.StatusBadRequest, err.Error())
    return
  }

  id :=   ctx.Query("id")
  if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

  excursion := schemas.Excursion{}

  if err := db.First(&excursion, id).Error; err != nil {
    sendError(ctx, http.StatusInternalServerError, "excursion not found")
    return
  }

  if req.Image != "" {
    excursion.Image = req.Image
  }
  if req.Title != "" {
    excursion.Title = req.Title
  }
  if req.Description != "" {
    excursion.Description = req.Description
  }
  if req.Buy != "" {
    excursion.Buy = req.Buy
  }
  if req.FindMore != "" {
    excursion.FindMore = req.FindMore
  }

  if err := db.Save(&excursion).Error; err != nil {
    logger.Errorf("error updating excursion %v", err.Error())
    sendError(ctx, http.StatusInternalServerError, "error updating excursion")
    return
  }

  ctx.JSON(http.StatusOK, gin.H {
    "updated data": excursion,
  })
}



func InitalizeHandler() {
  logger = config.GetLogger("handler")
  db = config.GetSQLite()
}