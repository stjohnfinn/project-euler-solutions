package main

import (
	"fmt"
	"math"
)

func sumDigits(x int) int {
	var sum int = 0

	for {
		if x < 10 {
			break
		}
		sum += x % 10
		x /= 10
	}
	sum += x
	return sum
}

func main() {
	for i := 0; i < 30; i++ {
		var curr int = int(math.Pow(2, float64(i)))
		fmt.Println("(", i, ",", sumDigits(curr), ")")
	}
}
