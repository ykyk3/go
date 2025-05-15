package statistics

import (
	"fmt"
	"math"

	"modulelearning/pkg/calculator"
)

// Average calculates the average of a slice of integers
func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0
	for _, num := range numbers {
		sum = calculator.Add(sum, num)
	}

	return float64(sum) / float64(len(numbers))
}

// StandardDeviation calculates the standard deviation of a slice of integers
func StandardDeviation(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}

	avg := Average(numbers)
	var sumSquaredDiff float64 = 0

	for _, num := range numbers {
		diff := float64(num) - avg
		sumSquaredDiff += diff * diff
	}

	return math.Sqrt(sumSquaredDiff / float64(len(numbers)))
}

// SumOfSquares calculates the sum of squares using the calculator package
func SumOfSquares(numbers []int) int {
	result := 0
	for _, num := range numbers {
		squared := calculator.Multiply(num, num)
		result = calculator.Add(result, squared)
	}
	return result
}

// Range calculates the range (max - min) of a slice of integers
func Range(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("cannot calculate range of empty slice")
	}

	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return calculator.Subtract(max, min), nil
} 