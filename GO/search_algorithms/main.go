package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const searchTimeDir = "search_execution_time/fedora/"

var testSizes = []int{10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

func generateTestData(size int) ([]int, int) {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr, arr[size-1]
}

func linearSearch(arr []int, targetValue int) int {
	for idx, value := range arr {
		if value == targetValue {
			return idx
		}
	}
	return -1
}

func binarySearch(arr []int, targetValue int) int {
	begin, end := 0, len(arr)-1
	for begin <= end {
		mid := begin + (end-begin)/2
		if arr[mid] == targetValue {
			return mid
		} else if arr[mid] < targetValue {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func countAlgorithmExecutionTime(arr []int, size int, targetValue int, funcName string) {
	start := time.Now()

	switch funcName {
	case "binary_search":
		binarySearch(arr, targetValue)
	default:
		linearSearch(arr, targetValue)
	}

	elapsed := time.Since(start).Seconds() * 1000 // Konwersja na milisekundy

	file, err := os.OpenFile(fmt.Sprintf("%s%s.csv", searchTimeDir, funcName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Cannot open file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{fmt.Sprintf("%d", size), fmt.Sprintf("%.6f", elapsed)}); err != nil {
		fmt.Println("Cannot write to file:", err)
	}
}

func performTestCases() {
	for _, size := range testSizes {
		arr, targetValue := generateTestData(size)
		countAlgorithmExecutionTime(arr, size, targetValue, "linear_search")
		countAlgorithmExecutionTime(arr, size, targetValue, "binary_search")
	}
}

func main() {
	performTestCases()
}
