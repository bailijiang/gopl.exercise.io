package tempconv

import "fmt"

// Celsius processes temperatures in the Celsius scale.
type Celsius float64

// Fahrenheit processes temperatures in the Fahrenheit scale.
type Fahrenheit float64

// Kelvin processes temperatures in the Kelvin scale.
type Kelvin float64

const (
	// AbsoluteZeroC is absolute zero Celsius.
	absoluteZeroC Celsius = -273.15
	// FreezingC is freezing Celsius.
	FreezingC Celsius = 0
	// ZeroKelvin is zero Kelvin.
	zeroKelvin Kelvin = -273.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - absoluteZeroC) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k + zeroKelvin) }
