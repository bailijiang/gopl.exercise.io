package main

import (
	"fmt"
	"os"
	"strings"
)

// Modify the echo program to also print os.Args[0], tha name of the command that invoked it.
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
