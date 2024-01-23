package services

import (
	"fmt"

	"github.com/mikeconroy/paceconv/types"
)

func ConvertDistanceToPace(req types.DistanceToPaceRequest) []types.Pace {
	var res []types.Pace

	distance := req.Distance
	// Convert distance to miles and then handle all conversions from miles
	if req.DistanceUnits == "kilometres" {
		distance = (KmToMiles(req.Distance))
	}

	durInMins := req.Duration.Minutes()
	minutesPerMile := types.Pace{
		Unit:         "mile",
		ReadableTime: getReadableMinutes(durInMins / distance),
	}
	minutesPerKm := types.Pace{
		Unit:         "kilometre",
		ReadableTime: getReadableMinutes(durInMins / MilesToKm(distance)),
	}
	res = append(res, minutesPerMile)
	res = append(res, minutesPerKm)
	return res
}

func MilesToKm(distance float64) float64 {
	return distance * 1.609344
}

func KmToMiles(distance float64) float64 {
	return distance / 1.609344
}

func getReadableMinutes(mins float64) string {
	wholeMins := int(mins)
	rem := mins - float64(wholeMins)
	secs := int(60 * rem)
	res := fmt.Sprintf("%d:%02d", wholeMins, secs)
	return res
}

func CelsiusToFahrenheit(temp float64) float64 {
	return (temp * (float64(9) / float64(5))) + 32
}

func FahrenheitToCelsius(temp float64) float64 {
	return (temp - 32) * (float64(5) / float64(9))
}
