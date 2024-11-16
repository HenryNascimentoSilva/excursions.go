package router

import (

	handler "excursion.com/handlers"
	"github.com/gin-gonic/gin"
)

func initializeroutes(router *gin.Engine) {
  // Initialize Handler
  handler.InitalizeHandler()
  v1 := router.Group("/api/v1")
  {
    v1.GET("/excursion/:id", handler.ShowExcursionHandler)
    v1.GET("/excursions", handler.ShowExcursionsHandler)
    v1.POST("/excursion", handler.CreateExcursionHandler)
    v1.DELETE("/excursion/:id", handler.DeleteExcursionHandler)
    v1.PUT("/excursion/:id", handler.UpdateExcursionHandler)
  }
}
