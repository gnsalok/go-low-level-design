package main

import "math"

type PriceCalculator interface {
	CalculateFare(trip *Trip) float64
}

type StandardPricingCalculator struct{}

const (
	baseFare      = 2.50
	costPerMile   = 1.50
	costPerMinute = 0.25
)

func (StandardPricingCalculator) CalculateFare(trip *Trip) float64 {
	distance := calculateDistance(trip.PickupLocation, trip.DropoffLocation)
	duration := estimateDuration(distance)
	return baseFare + (distance * costPerMile) + (duration * costPerMinute)
}

func calculateDistance(from, to *Location) float64 {
	latDiff := to.Latitude - from.Latitude
	lonDiff := to.Longitude - from.Longitude

	// it uses the Haversine formula (simplified)
	return math.Sqrt(math.Pow(latDiff, 2)+math.Pow(lonDiff, 2)) * 69
}

func estimateDuration(distanceMiles float64) float64 {
	return (distanceMiles / 30.0) * 60
}
