package main

import (
  "github.com/gin-gonic/gin"
  logger "github.com/sirupsen/logrus"
)

// @title Vehicle Management Payment Service
// @version 1.0
// @description Vehicle Management Service.
// @termsOfService http://swagger.io/terms/
//
// @contact.name API Support
// @contact.url https://github.com/kalilventura/vehicle-management-payment
// @contact.email kalilventur@gmail.com
//
// @license.name MIT License
// @license.url https://opensource.org/license/mit
func main() {
  defer handlePanic()

  engine := gin.Default()
  panic(engine.Run(":8080"))
}

func handlePanic() {
  if r := recover(); r != nil {
    logger.WithField("panic", r).
      Fatal("🚨 A critical and unrecoverable error occurred, forcing the application to stop.")
  }
}
