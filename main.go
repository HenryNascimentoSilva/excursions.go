package main

import (


	"excursion.com/config"
	"excursion.com/router"
)

var (
  logger *config.Logger
)

func main() {
  logger = config.GetLogger("main")

  err := config.Init()

  // Error Treatment
  if err != nil {
    logger.Errorf("Database Connection Error: %v", err)
    return
  }

  // Initialize Router
  router.Initialize()
}
