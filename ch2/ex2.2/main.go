// Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguments or from the standard input if there
// are no arguments, and convers each number into units like temperature in
// Celsius and Fahrenheit, length in feet and meters, weight in pounds and
// kilograms, and the like.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mttrs/gopl.io/ch2/ex2.1/tempconv"
)

func main() {
	// Read numbers from its command-line or from the standard input if ther is no
	// arguments.
	if len(os.Args[1:]) < 1 {
		fmt.Println("Input a number.")
		var t float64
		_, err := fmt.Scanf("%f", &t)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printConvertedValue(t)
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			printConvertedValue(t)
		}
	}
}

func printConvertedValue(t float64) {
	// Fahrenheit <=> Celsius
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))

	// Feet <=> Meter
	ft := tempconv.Feet(t)
	m := tempconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n",
		ft, tempconv.FtToM(ft), m, tempconv.MToFt(m))

	// Pound <=> Kilogram
	lb := tempconv.Pound(t)
	kg := tempconv.Kilogram(t)
	fmt.Printf("%s = %s, %s = %s\n",
		lb, tempconv.LbToKg(lb), kg, tempconv.KgToLb(kg))
}
