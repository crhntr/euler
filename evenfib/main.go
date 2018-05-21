/*
evenfib is an implementation of Problem 2: Even Fibonacci numbers.

Each new term in the Fibonacci sequence is generated by adding the
previous two terms. By starting with 1 and 2, the first 10 terms
will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values
do not exceed four million, find the sum of the even-valued terms.

*/
package main

func main() {
	var (
		max  = 4 * 1000000
		sum  int
		f, i int

		m1 = [2][2]int{{1, 1}, {1, 0}}
		m2 = [2][2]int{{1, 1}, {1, 0}}
	)

	for {

		m1[0][0] = m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0]
		m1[0][1] = m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1]
		m1[1][0] = m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0]
		m1[1][1] = m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1]

		if f = m1[0][1]; f >= max {
			break
		}

		if f%2 == 0 {
			sum += f
		}

		i++

	}

	println(sum)
}
