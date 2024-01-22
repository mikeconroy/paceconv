package types

type Pace struct {
	Unit         string
	ReadableTime string
}
type DistanceToPaceRequest struct {
	Distance float64
	Units    string
	Duration float64
}
