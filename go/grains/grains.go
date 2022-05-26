package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number <= 0 {
		return 0, errors.New("number must me greater than 0")
	} else if number > 64 {
		return 0, errors.New("a chessboard has no mmore than 64 squares")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	var totalGrains uint64 = 0

	for i := 1; i <= 64; i++ {
		grains, _ := Square(i)
		totalGrains += grains
	}

	return totalGrains
}
