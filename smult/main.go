/*
smult is a solution to euler project problem 5.

2520 is the smallest number that can be divided by each of the
numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly
divisible by all of the numbers from 1 to 20?
*/
package main

func main() {
	n := 20

nextN:
	for ; ; n++ {

		for i := 1; i < 20; i++ {
			if n%i != 0 {
				continue nextN
			}
		}

		println(n) // 232792560

		break
	}
}
