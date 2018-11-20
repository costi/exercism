package grains

import "errors"

// Square returns the number of grains on a specific square
// if every square we double the number of grains
func Square(position int) (uint64, error) {
	if position <= 0 {
		return 0, errors.New("Can't let you do that. No negativity allowed")
	}
	if position > 64 {
		return 0, errors.New("Can't go above 64")
	}
	return 1 << (uint(position) - 1), nil
}

// Total returns the total number of grains on a chess board
// if every square we double the number of grains
func Total() uint64 {
	return totalWithMax()
	// return totalWithSum()
	// return totalWithOr()
	// return totalWithOrInline()
}

// 0 0 0 0 1 +
// 0 0 0 1 0 +
// 0 0 1 0 0 +
// 0 1 0 0 0 +
// 1 0 0 0 0 =
// 1 1 1 1 1
func totalWithMax() uint64 {
	return ^uint64(0)
}

// more readable: sum += Square(i)
func totalWithSum() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		v, _ := Square(i)
		sum += v
	}
	return sum
}

// 0 0 0 0 1 |
// 0 0 0 1 0 |
// 0 0 1 0 0 |
// 0 1 0 0 0 |
// 1 0 0 0 0 =
// 1 1 1 1 1
func totalWithOr() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		v, _ := Square(i)
		sum = sum | v
	}
	return sum
}

// same as OR but inlining Square()
func totalWithOrInline() uint64 {
	var sum uint64
	for i := uint(0); i < 64; i++ {
		sum += (1 << i)
	}
	return sum
}
