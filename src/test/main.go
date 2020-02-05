package main

import (
	"fmt"
	"math"
)

func calTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {
	fmt.Println(calTriangle(3, 4))
}
