package functions

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius converts ºK to ºC
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

// same as the function above, but written as a method
// remade in temp-methods.go
// func (c celsius) kelvin() kelvin {
// 	return kelvin(c + 273.15)
// }

func TypeTemp() {
	var c celsius = 127.0
	var k kelvin
	k = celsiusToKelvin(c)
	fmt.Println(k)
	// k = c.kelvin()
	fmt.Print(c, "º C is ", k, "º K")
}
