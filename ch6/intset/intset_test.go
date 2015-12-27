package intset

import (
	"fmt"
	"testing"

	"github.com/mttrs/gopl.io/ch6/intset"
)

func TestHas(t *testing.T) {
	var x intset.IntSet
	x.Add(0)
	x.Add(64)
	tests := []struct {
		input int
		want  bool
	}{
		{0, true},
		{64, true},
		{3, false},
		{128, false},
	}
	for _, test := range tests {
		if got := x.Has(test.input); got != test.want {
			t.Errorf("Has(%d) = %t", test.input, got)
		}
	}
}

// May be bad. It looks like String() method test...
func TestAdd(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{0, "{0}"},
		{1, "{1}"},
		{64, "{64}"},
	}
	for _, test := range tests {
		var x intset.IntSet
		x.Add(test.input)
		if got := x.String(); got != test.want {
			t.Errorf("Add(%d) = %s", test.input, got)
		}
	}
}

func TestSingleAddLen(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{0, 1},
		{1, 1},
		{64, 1},
		{256, 1},
	}

	// single add value case
	for _, test := range tests {
		var x intset.IntSet
		x.Add(test.input)
		if got := x.Len(); got != test.want {
			t.Errorf("Len() input %d = %d", test.input, got)
		}
	}
}

func TestMultiAddLen(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0, 1}, 2},
		{[]int{0, 0}, 1}, // duplicated input
		{[]int{0, 1, 2}, 3},
		{[]int{0, 64, 128}, 3},
	}

	for _, test := range tests {
		var x intset.IntSet
		for _, v := range test.input {
			x.Add(v)
		}
		if got := x.Len(); got != test.want {
			t.Errorf("Len() input %v = %d", test.input, got)
		}
	}
}

func TestElems(t *testing.T) {}

func Example_one() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func Example_ex6_1() {
	var x, z intset.IntSet
	x.Add(0)
	x.Add(128)
	x.Add(1)
	x.Add(2)
	x.Add(37)
	x.Add(64)

	fmt.Println(x.String())

	// Show the number of x elements
	fmt.Println(x.Len())

	x.Remove(0)
	x.Remove(64)
	x.Remove(128)
	x.Remove(2)
	fmt.Println(x.String())

	z = x.Copy()
	fmt.Println(z.String())

	// Confirm that z doesn't have x's reference.
	x.Add(5)
	fmt.Println(x.String())
	fmt.Println(z.String())

	// Clear x
	x.Clear()
	fmt.Println(x.String())

	// Output:
	// {0 1 2 37 64 128}
	// 6
	// {1 37}
	// {1 37}
	// {1 5 37}
	// {1 37}
	// {}
}

func Example_ex6_2() {
	var x intset.IntSet
	x.AddAll(0, 1, 2, 64, 128, 3)
	fmt.Println(x.String())

	// Output:
	// {0 1 2 3 64 128}
}

func Example_ex6_4() {
	var x intset.IntSet
	x.AddAll(0, 1, 64, 128, 5)
	fmt.Println(x.Elems())

	// Output:
	// [0 1 5 64 128]
}
