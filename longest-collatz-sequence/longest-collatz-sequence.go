package main

import (
	"fmt"
)

var g_sols = make(map[int]int)

func main() {
	var seqLengths []int
	for i := 1; i < 1000000; i++ {
		seqLengths = append(seqLengths, collatz(i))
	}

	max := 0
	max_index := 0

	for i := 0; i < 1000000; i++ {
		if g_sols[i] > max {
			max = g_sols[i]
			max_index = i
		}
	}

	fmt.Println("Longest collatz sequence: ", max_index, " : ", max)
}

func collatz(n int) int {

	var l int = 1

	var temp int = n

	fmt.Println("Calculating collatz(", n, ")")

	for {
		if temp == 1 {
			break
		}
		if g_sols[temp] != 0 {
			fmt.Println("using existing solution: ")
			fmt.Println("map [", temp, " : ", g_sols[temp], "]")
			l += g_sols[temp] - 1
			break
		}
		if temp%2 == 0 {
			temp /= 2
		} else {
			temp = (3 * temp) + 1
		}
		l++
	}

	g_sols[n] = l

	return l
}
