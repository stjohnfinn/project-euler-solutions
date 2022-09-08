package main

import (
	"fmt"
)

func Binomial(n, k int) int {
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

func main() {
	fmt.Println(Binomial(40, 20))
}
