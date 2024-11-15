package router

import (

	handler "excursion.com/handlers"
	"github.com/gin-gonic/gin"
)

func initializeroutes(router *gin.Engine) {
  v1 := router.Group("/api/v1")
  {
    v1.GET("/excursion", handler.ShowExcursionHandler)
    v1.POST("/excursion", handler.CreateExcursionHandler)
    v1.DELETE("/excursion", handler.DeleteExcursionHandler)
    v1.PUT("/excursion", handler.UpdateExcursionHandler)
    v1.GET("/excursion", handler.ShowExcursionsHandler)
  }
}
