package tempconv

import "fmt"

// Celsius processes temperatures in the Celsius scale.
type Celsius float64

// Fahrenheit processes temperatures in the Fahrenheit scale.
type Fahrenheit float64

// Kelvin processes temperatures in the Kelvin scale.
type Kelvin float64

// Meter is a unit of length of the path travelled by light in a vacuum
// during a time interval of 1/299,792,458 of a second.
type Meter float64

// Feet is a unit of length used in the imperial and U.S. customary measurement
// systems, representing 1/3 of a yard, and is subdivided into twelve inches.
type Feet float64

// Kilogram is a unit of weight
type Kilogram float64

// Pound is a unit of weight
type Pound float64

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

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string { return fmt.Sprintf("%gft", ft) }

func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
func (lb Pound) String() string    { return fmt.Sprintf("%glb", lb) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - absoluteZeroC) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k + zeroKelvin) }

// MToFt converts meter to feet.
func MToFt(m Meter) Feet { return Feet(m * 3.2808) }

// FtToM converts meter to feet.
func FtToM(ft Feet) Meter { return Meter(ft / 3.2808) }

// KgToLb converts Kilogram to pound.
func KgToLb(kg Kilogram) Pound { return Pound(kg * 2.2046) }

// LbToKg converts pound to Kilogram.
func LbToKg(lb Pound) Kilogram { return Kilogram(lb / 2.2046) }
