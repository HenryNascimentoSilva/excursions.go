package handler

import (

	"github.com/gin-gonic/gin"
)

func CreateExcursionHandler(ctx *gin.Context ) {
	request := CreateExcursionRequest{}

  ctx.BindJSON(&request)

  if err := request.Validate(); err!=nil {
    logger.Errorf("validation error: %v", err.Error())
    return
  }

  if err := db.Create(&request).Error; err!=nil {
    logger.Errorf("Error Creating Excursion: %v", err.Error())
    return
  }
}