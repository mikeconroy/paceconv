package types

import "time"

type Pace struct {
	Unit         string
	ReadableTime string
}
type DistanceToPaceRequest struct {
	Distance      float64
	DistanceUnits string
	Duration      time.Duration
}

type Distance struct {
	Value float64
	Units string
}

type Temperature struct {
	Value float64
	Units string
}
