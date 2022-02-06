package main

import "fmt"

func main() {
	squareOfSum := 0
	sumOfSquare := 0
	for i := 0; i <= 100; i++ {
		squareOfSum += i
	}
	for j := 1; j <= 100; j++ {
		sumOfSquare = sumOfSquare + (j * j)
	}
	squareOfSum = squareOfSum * squareOfSum
	answer := squareOfSum - sumOfSquare
	fmt.Println(answer)
}
