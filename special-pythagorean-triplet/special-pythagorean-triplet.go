package main

import (
	"fmt"
	"math/rand"
)

func main() {
	findTripletFromTotal(1000)
}

func findTripletFromTotal(target int) {
	for {
		a := rand.Intn(target-1) + 1
		b := rand.Intn(target-1) + 1
		c := rand.Intn(target-1) + 1

		if a*a+b*b == c*c {
			fmt.Println("Pythagorean Triple Found!")
			fmt.Println(a, " ", b, " ", c)

			if a+b+c == target {
				fmt.Println("Solution Found")
				fmt.Println("a + b + c = ", a+b+c)
				fmt.Println("a * b * c = ", a*b*c)
				break
			}
		}
	}
}
