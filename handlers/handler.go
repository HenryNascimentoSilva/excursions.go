package handler

import (
	"net/http"

	"excursion.com/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
  logger *config.Logger
  db *gorm.DB
)

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

func InitalizeHandler() {
  logger = config.GetLogger("handler")
  db = config.GetSQLite()
}