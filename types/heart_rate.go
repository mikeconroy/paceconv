package types

type HeartRateZones struct {
	MaxHr int
	Zones []HeartRateZone
}

type HeartRateZone struct {
	ZoneNum   int
	Intensity string
	Type      string
	MinHr     int
	MaxHr     int
	MinEffort int
	MaxEffort int
}
