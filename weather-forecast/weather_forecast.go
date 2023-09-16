// Package weather provides functionality for managing weather conditions.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current location for which weather is being tracked.
var CurrentLocation string

// Forecast updates the current location and condition and returns a weather forecast.
// It takes two parameters: city (the location) and condition (the weather condition).
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition

	forecast := CurrentLocation + " - current weather condition: " + CurrentCondition

	return forecast
}
