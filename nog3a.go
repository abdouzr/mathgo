package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// Average calculates the mean of a slice of float64 numbers
func Average(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

// Median calculates the median of a slice of float64 numbers
func Median(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sort.Float64s(nums)
	n := len(nums)
	if n%2 == 0 {
		return (nums[n/2-1] + nums[n/2]) / 2
	}
	return nums[n/2]
}

// Variance calculates the variance of a slice of float64 numbers
func Variance(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	mean := Average(nums)
	var sumSquares float64
	for _, num := range nums {
		diff := num - mean
		sumSquares += diff * diff
	}
	return sumSquares / float64(len(nums))
}

// StandardDeviation calculates the standard deviation of a slice of float64 numbers
func StandardDeviation(nums []float64) float64 {
	return math.Sqrt(Variance(nums))
}

// ReadFile reads numbers from a file and returns them as a slice of float64
func ReadFile(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}

	filePath := os.Args[1]
	numbers, err := ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	avg := Average(numbers)
	med := Median(numbers)
	variance := Variance(numbers)
	stdDev := StandardDeviation(numbers)

	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	fmt.Printf("Median: %d\n", int(math.Round(med)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}

