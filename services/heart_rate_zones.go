package services

import "github.com/mikeconroy/paceconv/types"

// 68% - 73% - Easy 	- Recovery
// 73% - 80% - Steady	- Aerobic
// 80% - 87% - Moderate - Tempo
// 87% - 93% - Hard 	- Lactate Threshold
// 93% - 100% - Maximum - Anaerobic

var zonesToCalculate = [5]types.HeartRateZone{
	{
		ZoneNum:   1,
		Intensity: "Easy",
		Type:      "Recovery",
		MinEffort: 55,
		MaxEffort: 65,
	}, {
		ZoneNum:   2,
		Intensity: "Steady",
		Type:      "Aerobic",
		MinEffort: 66,
		MaxEffort: 75,
	}, {
		ZoneNum:   3,
		Intensity: "Moderate",
		Type:      "Tempo",
		MinEffort: 76,
		MaxEffort: 85,
	}, {
		ZoneNum:   4,
		Intensity: "Hard",
		Type:      "Lactate Threshold",
		MinEffort: 86,
		MaxEffort: 90,
	}, {
		ZoneNum:   5,
		Intensity: "Maximum",
		Type:      "Anaerobic",
		MinEffort: 91,
		MaxEffort: 100,
	},
}

func CalculateHeartRateZones(maxHr int) types.HeartRateZones {
	var zones []types.HeartRateZone

	for _, zone := range zonesToCalculate {
		zone.MinHr = int((float64(zone.MinEffort) / 100) * float64(maxHr))
		zone.MaxHr = int((float64(zone.MaxEffort) / 100) * float64(maxHr))
		zones = append(zones, zone)
	}

	return types.HeartRateZones{
		MaxHr: maxHr,
		Zones: zones,
	}
}
