// Package weather This is the package to handle our weather app.
package weather

var (
	// CurrentCondition This holds the current weather condition for the specified location.
	CurrentCondition string
	// CurrentLocation This holds the exact location of the place at the time the code is being run.
	CurrentLocation string
)

// Forecast This is the method that returns the weather condition for the current location the president is.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
