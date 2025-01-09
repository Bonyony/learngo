package functions

import "fmt"

// type celsius float64
// type kelvin float64
type fahrenheit float64

// fahrenheit to celsius equation
// f := (c * 9.0 / 5.0) + 32.0

// celsius to kelvin
func (c celsius) c2kelvin() kelvin {
	return kelvin(c + 273.15)
}

// celsius to fahrenheit
func (c celsius) c2fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// kelvin to f
func (k kelvin) k2fahrenheit() fahrenheit {
	return k.k2celsius().c2fahrenheit()
}

//  kelvin to celcius
func (k kelvin) k2celsius() celsius {
	return celsius(k - 273.15)
}

// fahr to cel
func (f fahrenheit) f2celsius() celsius {
	return celsius((f - 32.0) * (5.0 / 9.0))
}

// fahr to kelv
func (f fahrenheit) f2kelvin() kelvin {
	return f.f2celsius().c2kelvin()
}

func TempMethods() {
	var c celsius = 0.0
	fmt.Println(c.c2fahrenheit())
}
