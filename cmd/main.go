package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mikeconroy/paceconv/handler"
)

func main() {
	e := echo.New()
	indexHandler := &handler.IndexHandler{}
	conversionHandler := &handler.ConversionHandler{}
	e.GET("/", indexHandler.HandleIndex)
	e.POST("/convert/distance-to-pace", conversionHandler.HandleDistanceToPace)
	e.POST("/convert/distance", conversionHandler.HandleDistanceConversion)
	e.POST("/convert/temperature", conversionHandler.HandleTemperatureConversion)
	e.POST("/convert/heart-rate-zones", conversionHandler.HandleHeartRateZones)
	e.Static("/static", "assets")
	e.Logger.Fatal(e.Start(":3000"))
}
