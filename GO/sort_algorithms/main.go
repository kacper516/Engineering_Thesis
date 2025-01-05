package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	sort_test_cases_file = "/home/kacper516/engineering-thesis/GO/sort_algorithms/test.csv"
	sort_time_dir        = "sort_execution_time/fedora/"
)

func bubble_sort(arr []int, size int) {
	for i := 0; i < size; i++ {
		swapped := false
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !swapped {
			return
		}
	}
}

func partition(arr []int, lower int, upper int) int {
	pivot := arr[upper]
	i := lower - 1

	for j := lower; j < upper; j++ {
		if arr[j] <= pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[i+1], arr[upper] = arr[upper], arr[i+1]
	return i + 1
}

func quick_sort(arr []int, lower int, upper int) {
	if lower < upper {
		pivot := partition(arr, lower, upper)
		quick_sort(arr, lower, pivot-1)
		quick_sort(arr, pivot+1, upper)
	}
}

func heapify(arr []int, current_index int, size int) {
	current_max_index := current_index
	left_node := current_max_index*2 + 1
	right_node := current_max_index*2 + 2

	if left_node < size && arr[left_node] > arr[current_max_index] {
		current_max_index = left_node
	}

	if right_node < size && arr[right_node] > arr[current_max_index] {
		current_max_index = right_node
	}

	if current_index != current_max_index {
		arr[current_index], arr[current_max_index] = arr[current_max_index], arr[current_index]
		heapify(arr, current_max_index, size)
	}
}

func heap_sort(arr []int, size int) {
	for i := size/2 - 1; i >= 0; i-- {
		heapify(arr, i, size)
	}

	for i := size - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func count_algorithm_execution_time(arr []int, size int, func_name string) {
	start := time.Now()
	switch func_name {
	case "quick_sort":
		quick_sort(arr, 0, size-1)
	case "heap_sort":
		heap_sort(arr, size)
	default:
		bubble_sort(arr, size)
	}
	elapsed := time.Since(start).Seconds() * 1000.0

	file, err := os.OpenFile(fmt.Sprintf("%s%s.csv", sort_time_dir, func_name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Cannot open file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{strconv.Itoa(size), fmt.Sprintf("%.6f", elapsed)})
	if err != nil {
		fmt.Println("Cannot write to file:", err)
	}
}

func perform_test_cases() {
	file_contents, err := os.ReadFile(sort_test_cases_file)
	if err != nil {
		fmt.Println("Error while reading file:", err)
		return
	}
	lines := strings.Split(string(file_contents), "\n")

	for i := 0; i < len(lines); i += 2 {
		if i+1 >= len(lines) {
			break
		}
		size, err := strconv.Atoi(lines[i])
		if err != nil {
			fmt.Println("Error parsing size:", err)
			continue
		}
		arrStr := strings.Split(lines[i+1], ",")
		arr := make([]int, size)
		for j, s := range arrStr {
			arr[j], err = strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error parsing array element:", err)
				continue
			}
		}

		arrBubble := make([]int, size)
		copy(arrBubble, arr)
		count_algorithm_execution_time(arrBubble, size, "bubble_sort")

		arrQuick := make([]int, size)
		copy(arrQuick, arr)
		count_algorithm_execution_time(arrQuick, size, "quick_sort")

		arrHeap := make([]int, size)
		copy(arrHeap, arr)
		count_algorithm_execution_time(arrHeap, size, "heap_sort")
	}
}

func main() {
	perform_test_cases()
}
