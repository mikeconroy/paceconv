package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikeconroy/paceconv/services"
	"github.com/mikeconroy/paceconv/types"
	"github.com/mikeconroy/paceconv/view/results"
)

type ConversionHandler struct{}

func (h *ConversionHandler) HandleDistanceToPace(c echo.Context) error {
	fmt.Println("Handling Distance to Pace Conversion")
	distance := c.FormValue("distance")
	distanceUnits := c.FormValue("units")
	duration := c.FormValue("duration")
	fmt.Println(distance, distanceUnits, duration)
	distanceFloat, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		// Change this to return the render of an error
		return fmt.Errorf("Unable to convert distance to Float")
	}

	dur, err := time.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("Unable to convert duration to time.Duration")
	}

	req := types.DistanceToPaceRequest{
		Distance:      distanceFloat,
		DistanceUnits: distanceUnits,
		Duration:      dur,
	}

	res := services.ConvertDistanceToPace(req)
	segment := results.ShowDistanceToPace(req, res)
	return segment.Render(c.Request().Context(), c.Response())
}

func (h *ConversionHandler) HandleDistanceConversion(c echo.Context) error {
	fmt.Println("Handling Distance conversion")
	distance := c.FormValue("distance")
	distanceUnits := c.FormValue("units")
	fmt.Println(distance, distanceUnits)
	distanceFloat, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		return fmt.Errorf("Unable to convert distance to Float")
	}

	req := types.Distance{
		Value: distanceFloat,
		Units: distanceUnits,
	}
	res := types.Distance{}

	if req.Units == "miles" {
		res.Units = "kilometres"
		res.Value = services.MilesToKm(req.Value)
	} else if req.Units == "kilometres" {
		res.Units = "miles"
		res.Value = services.KmToMiles(req.Value)
	} else {
		return fmt.Errorf("Unrecognised distance units")
	}

	segment := results.ShowDistanceConversion(req, res)
	return segment.Render(c.Request().Context(), c.Response())
}

func (h *ConversionHandler) HandleTemperatureConversion(c echo.Context) error {
	fmt.Println("Handling Temperature Conversion")
	temp := c.FormValue("temperature")
	tempUnits := c.FormValue("units")
	fmt.Println(temp, tempUnits)
	tempFloat, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return fmt.Errorf("Unable to convert Temperature to Float.")
	}

	req := types.Temperature{
		Value: tempFloat,
		Units: tempUnits,
	}
	res := types.Temperature{}

	if req.Units == "celsius" {
		res.Value = services.CelsiusToFahrenheit(req.Value)
		res.Units = "fahrenheit"
	} else if req.Units == "fahrenheit" {
		res.Value = services.FahrenheitToCelsius(req.Value)
		res.Units = "celsius"
	} else {
		return fmt.Errorf("Unrecognised temperature unit.")
	}

	segment := results.ShowTemperatureConversion(req, res)
	return segment.Render(c.Request().Context(), c.Response())
}

func (h *ConversionHandler) HandleHeartRateZones(c echo.Context) error {
	fmt.Println("Handling Heart Rate Zones")
	maxHrStr := c.FormValue("maxHeartRate")
	fmt.Println(maxHrStr)
	maxHr, err := strconv.Atoi(maxHrStr)
	if err != nil {
		return fmt.Errorf("Unable to convert maxHr to Int.")
	}
	res := services.CalculateHeartRateZones(maxHr)

	segment := results.ShowHRZones(res)
	return segment.Render(c.Request().Context(), c.Response())
}
