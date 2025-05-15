package main

import (
	"fmt"
	"log"

	"modulelearning/pkg/calculator"
	"modulelearning/pkg/statistics"
)

func main() {
	// calculator パッケージのテスト
	fmt.Println("===== Calculator Package Tests =====")
	
	// 足し算のテスト
	result := calculator.Add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// 引き算のテスト
	result = calculator.Subtract(10, 4)
	fmt.Printf("10 - 4 = %d\n", result)

	// 掛け算のテスト
	result = calculator.Multiply(6, 7)
	fmt.Printf("6 * 7 = %d\n", result)

	// 割り算のテスト
	result, err := calculator.Divide(20, 5)
	if err != nil {
		log.Fatalf("Error in division: %v", err)
	}
	fmt.Printf("20 / 5 = %d\n", result)

	// エラーケース: ゼロ除算
	_, err = calculator.Divide(10, 0)
	if err != nil {
		fmt.Printf("Division by zero error: %v\n", err)
	}

	// statistics パッケージのテスト
	fmt.Println("\n===== Statistics Package Tests =====")
	
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Printf("Numbers: %v\n", numbers)
	
	// 平均値の計算
	avg := statistics.Average(numbers)
	fmt.Printf("Average: %.2f\n", avg)
	
	// 標準偏差の計算
	stdDev := statistics.StandardDeviation(numbers)
	fmt.Printf("Standard Deviation: %.2f\n", stdDev)
	
	// 二乗和の計算
	sumOfSquares := statistics.SumOfSquares(numbers)
	fmt.Printf("Sum of Squares: %d\n", sumOfSquares)
	
	// 範囲の計算
	rng, err := statistics.Range(numbers)
	if err != nil {
		log.Fatalf("Error calculating range: %v", err)
	}
	fmt.Printf("Range: %d\n", rng)
} 