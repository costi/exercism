package diffsquares

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return (sum * sum)
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the diff between square of sums and sum of squares in O(1)
func Difference(n int) int {
	return (SquareOfSum(n) - SumOfSquares(n))
}
