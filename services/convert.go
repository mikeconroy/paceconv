package services

import (
	"fmt"
	"time"

	"github.com/mikeconroy/paceconv/types"
)

func ConvertDistanceToPace(req types.DistanceToPaceRequest) []types.Pace {
	var res [8]types.Pace

	distance := req.Distance
	// Convert distance to miles and then handle all conversions from miles
	if req.DistanceUnits == "kilometres" {
		distance = (KmToMiles(req.Distance))
	}

	nanos := float64(req.Duration.Nanoseconds())
	minutesPerMile := types.Pace{
		Unit:         "1 mile",
		ReadableTime: getReadableMinutes(nanos / distance),
	}
	minutesPerKm := types.Pace{
		Unit:         "1 kilometre",
		ReadableTime: getReadableMinutes(nanos / MilesToKm(distance)),
	}
	minutesPer5Km := types.Pace{
		Unit:         "5 kilometres",
		ReadableTime: getReadableMinutes((nanos * 5) / MilesToKm(distance)),
	}
	minutesPer5Miles := types.Pace{
		Unit:         "5 miles",
		ReadableTime: getReadableMinutes((nanos * 5) / distance),
	}
	minutesPer10Km := types.Pace{
		Unit:         "10 kilometres",
		ReadableTime: getReadableMinutes((nanos * 10) / MilesToKm(distance)),
	}
	minutesPer10Miles := types.Pace{
		Unit:         "10 miles",
		ReadableTime: getReadableMinutes((nanos * 10) / distance),
	}
	minutesPerHM := types.Pace{
		Unit:         "Half Marathon - 13.1 miles - 21.08 kilometres",
		ReadableTime: getReadableMinutes((nanos * 13.1) / distance),
	}
	minutesPerMarathon := types.Pace{
		Unit:         "Marathon - 26.2 miles - 42.16 kilometres",
		ReadableTime: getReadableMinutes((nanos * 26.2) / distance),
	}
	res[0] = minutesPerMile
	res[1] = minutesPerKm
	res[2] = minutesPer5Km
	res[3] = minutesPer5Miles
	res[4] = minutesPer10Km
	res[5] = minutesPer10Miles
	res[6] = minutesPerHM
	res[7] = minutesPerMarathon
	return res[:]
}

func MilesToKm(distance float64) float64 {
	return distance * 1.609344
}

func KmToMiles(distance float64) float64 {
	return distance / 1.609344
}

func getReadableMinutes(dur float64) string {
	durD := time.Duration(dur)
	hours := int(durD.Hours())
	wholeMins := int(durD.Minutes()) % 60
	secs := int(durD.Seconds()) % 60

	var res string
	if hours > 0 {
		res = fmt.Sprintf("%d:%02d:%02d", hours, wholeMins, secs)
	} else {
		res = fmt.Sprintf("%d:%02d", wholeMins, secs)
	}
	// fmt.Println("Input", mins, "hours", hours, "rem", rem, "secs", secs, "wholeMins", wholeMins)
	return res
}

func CelsiusToFahrenheit(temp float64) float64 {
	return (temp * (float64(9) / float64(5))) + 32
}

func FahrenheitToCelsius(temp float64) float64 {
	return (temp - 32) * (float64(5) / float64(9))
}
