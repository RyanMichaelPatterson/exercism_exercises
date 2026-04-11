// Package weather provides tools to display
// the weather in a location.
package weather

var (
    // CurrentCondition represents the current weather condition.
	CurrentCondition string
    // CurrentLocation reresents the current location.
	CurrentLocation  string
)

// Forecast returns a string value that contains the weather for a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
