// Package weather provides functionality to track and display
// current weather conditions in different locations.
package weather

// CurrentCondition stores the current weather condition (e.g., "sunny", "rainy").
var CurrentCondition string

// CurrentLocation stores the current location (e.g., "Frostpeak", "Bumblebee Bay").
var CurrentLocation string

// Forecast takes a city name and the current weather condition,
// updates the CurrentLocation and CurrentCondition variables,
// and returns a formatted string with the weather information.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
