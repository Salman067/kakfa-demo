package main

import (
	consumer "kafka/kafka"
	producer "kafka/kafka"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	api := app.Group("/api/v1")
	api.POST("/comment", producer.CreateComment)
	api.GET("/comment", consumer.GetComment)
	app.Logger.Fatal(app.Start(":3000"))
}
