package diffsquares

func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	squareofSum := sum * sum
	return squareofSum
}

func SumOfSquares(n int) int {
	add := (n * (n + 1) * ((n * 2) + 1)) / 6
	return add
}

func Difference(n int) int {
	difference := SquareOfSum(n) - SumOfSquares(n)
	return difference
}
