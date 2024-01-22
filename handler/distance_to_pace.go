package handler

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mikeconroy/paceconv/types"
	"github.com/mikeconroy/paceconv/view"
)

type ConversionHandler struct{}

func (h *ConversionHandler) HandleDistanceToPace(c echo.Context) error {
	fmt.Println("Handling Distance to Pace Conversion")
	distance := c.FormValue("distance")
	units := c.FormValue("units")
	duration := c.FormValue("duration")

	distanceFloat, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		// Change this to return the render of an error
		return fmt.Errorf("Unable to convert distance to int")
	}
	durationFloat, err := strconv.ParseFloat(duration, 64)
	if err != nil {
		return fmt.Errorf("Unable to convert duration to int")
	}

	req := types.DistanceToPaceRequest{
		Distance: distanceFloat,
		Units:    units,
		Duration: durationFloat,
	}
	fmt.Println(distance, units, duration)
	page := view.ShowDistanceToPace(req, convertDistanceToPace(distanceFloat, units, durationFloat))
	return page.Render(c.Request().Context(), c.Response())
}

func convertDistanceToPace(distance float64, unit string, duration float64) []types.Pace {
	var res []types.Pace
	// Convert distance to miles and then handle all conversions from miles
	if unit == "kilometres" {
		distance = (distance / 1.609344)
	}
	minutesPerMile := types.Pace{
		Unit:         "mile",
		ReadableTime: fmt.Sprint(duration / distance),
	}
	minutesPerKm := types.Pace{
		Unit:         "kilometre",
		ReadableTime: fmt.Sprint(duration / (distance * 1.609344)),
	}
	res = append(res, minutesPerMile)
	res = append(res, minutesPerKm)
	return res
}
