package service

import "math"

func CalculateAttempts(min, max int) (easy, medium, hard int) {
	numbersCount := max - min + 1

	easy = int(math.Ceil(math.Log2(float64(numbersCount))))
	medium = int(math.Ceil(float64(easy) * 0.75))
	hard = int(math.Ceil(float64(medium) * 0.75))

	return
}
