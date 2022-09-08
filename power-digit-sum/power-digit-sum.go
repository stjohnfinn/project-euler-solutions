package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {

	fmt.Println("Starting power-digit-sum.go...")
	x := big.NewInt(int64(math.Pow(2, 1000)))
	sum := 0
	for _, val := range strconv.Itoa(x) {
		sum += strconv.ParseInt(val)
	}
}
