/*
evenfib is an implementation of Problem 3: Largest prime factor.

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import "math"

func main() {
	var f int
	for f = range factorize(600851475143) {
	}
	println(f)
}

func factorize(n int) <-chan int {
	factors := make(chan int)

	go func() {
		for n%2 == 0 {
			factors <- 2
			n /= 2
		}

		sqrtn := int(math.Sqrt(float64(n)))
		for i := 3; i <= sqrtn; i += 2 {
			for n%i == 0 {
				factors <- i
				n /= i
			}
		}

		if n > 2 {
			factors <- n
		}
		close(factors)
	}()

	return factors
}
