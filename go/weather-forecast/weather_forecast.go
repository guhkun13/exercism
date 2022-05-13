// Package weather for checking weather.
package weather

// CurrentCondition is weather condition at the time.
var CurrentCondition string

// CurrentLocation is location at the time.
var CurrentLocation string

// Forecast weather by location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
