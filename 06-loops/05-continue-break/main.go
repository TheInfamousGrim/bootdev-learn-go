package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		if i == 2 {
			fmt.Println(i)
		}

		if i%2 == 0 {
			continue
		}

		sq_root := int(math.Sqrt(float64(i)))
		brokeLoop := false
		pBrokeLoop := &brokeLoop

		for j := 2; j <= sq_root; j++ {
			if i%j == 0 {
				*pBrokeLoop = true
				break
			}
		}

		if brokeLoop {
			continue
		}
		fmt.Println(i)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
