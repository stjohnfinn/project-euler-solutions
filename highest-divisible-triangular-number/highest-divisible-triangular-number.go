package main

import (
	"math"
	"fmt"
)

func solve() int {
	triangle := 0
	for i := 0; ; i++ {
		triangle += i
		fmt.Println("Checking ", triangle)
		if countDivisors(triangle) > 500 {
			fmt.Println("Solution found! -> ", triangle)
			return triangle
		}
	}
	return 0
}

func countDivisors(n int) int {
	count := 0
	end := int(math.Sqrt(float64(n)))
	for i := 1; i < end; i++ {
		if (n % i == 0) {
			count += 2
		}
	}
	if end * end == n {
		count++
	}
	return count
}

func main() {
	fmt.Println(solve())
}
