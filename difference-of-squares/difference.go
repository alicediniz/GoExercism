// Package diffsquares implements
package diffsquares

// SquareOfSum calculates the square of the sum of the first num natural numbers
// https://en.wikipedia.org/wiki/1_%2B_2_%2B_3_%2B_4_%2B_%E2%8B%AF
func SquareOfSum(num int) int {
	sum := num * (num + 1) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of the squares of the first num natural numbers
// https://en.wikipedia.org/wiki/Square_pyramidal_number
func SumOfSquares(num int) int {
	sum := num * (num + 1) * (2*num + 1) / 6
	return sum
}

// Difference gives the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
