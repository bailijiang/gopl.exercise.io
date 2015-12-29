package main

// Write const declarations for KB, MB, up through YB as compactly as you can.

// It looks like impossible to use iota to resolve the exercise.
// http://stackoverflow.com/questions/34124294/writing-powers-of-10-as-constants-compactly
const (
	KB, MB, GB, TB, PB, EB, ZB, YB = 1e3, 1e6, 1e9, 1e12, 1e15, 1e18, 1e21, 1e24
)

func main() {}
