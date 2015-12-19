// Add types, constants and functions to tempconv for processing the temperatures
// in the Kelvin scale, where zero Kelvin is -273.15°C and a difference of 1K has
// the same magnitude as 1°C.
package main

import (
	"fmt"

	"github.com/mttrs/gopl.io/ch2/ex2.1/tempconv"
)

func main() {
	c := tempconv.FreezingC
	fmt.Println("Freezing Celsius:", c)
	fmt.Println("C(0˚C) to F:", tempconv.CToF(c))
	fmt.Println("C(0˚C) to K:", tempconv.CToK(c))
	fmt.Println("C(20˚C) to F:", tempconv.CToF(c+20))
	fmt.Println("F(0˚F) to C:", tempconv.FToC(0))
	fmt.Println("F(32˚F) to C:", tempconv.FToC(32))
}
