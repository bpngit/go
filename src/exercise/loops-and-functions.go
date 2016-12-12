package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sqrt(9))
}

// Newton's formula to get sqrt.
const epsilon = .000000005 // a value which defines the threashold
const initial = 10.0

func Sqrt(x float64) (int, float64) {
	result := initial
	// formula  value = x - (x2 -x)/2x
	formulaValue := func() float64 {
		return result - (result*result-x)/(2*result)
	}
	iteration := 1
	for i := formulaValue(); math.Abs(i-result) > epsilon; {
		result = i
		i = formulaValue()
		iteration++

	}
	return iteration, result

}
