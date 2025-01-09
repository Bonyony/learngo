package functions

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}
func kelvinToFahrenheit(k float64) float64 {
	f := celsiusToFahrenheit(kelvinToCelsius(k))
	return f
}

func Temperature() {
	// this is an arbitrary value
	kelvin := 0.0

	// celsius := kelvinToCelsius(kelvin)
	// fahrenheit := celsiusToFahrenheit(celsius)
	// fmt.Println(kelvin, "º K is ", celsius, "º C")
	// fmt.Println(celsius, "º C is ", fahrenheit, "º F")

	fahrenheit := kelvinToFahrenheit(kelvin)
	fmt.Println(kelvin, "º K is ", fahrenheit, "º F")
}
