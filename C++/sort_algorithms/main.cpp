#include <chrono>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

const std::string sort_time = "C++/sort_algorithms/sort_execution_time/";
const std::string sort_test_cases_file =
    "C:/Users/Kacper/Documents/Engeenering_thesis/sort_test_cases.csv";

void bubble_sort(std::vector<int> &arr, int size) {
  for (int i = 0; i < size; i++) {
    bool swapped = false;
    for (int j = 0; j < size - i - 1; j++) {
      if (arr[j] > arr[j + 1]) {
        swapped = true;
        std::swap(arr[j + 1], arr[j]);
      }
    }
    if (!swapped) {
      return; // array is sorted
    }
  }
}

int partition(std::vector<int> &arr, int lower, int upper) {
  int pivot = arr[upper];
  int i = lower - 1;

  for (int j = lower; j < upper; j++) {
    if (arr[j] <= pivot) {
      i++;
      std::swap(arr[j], arr[i]);
    }
  }

  std::swap(arr[upper], arr[i + 1]);
  return i + 1;
}

void quick_sort(std::vector<int> &arr, int lower, int upper) {
  if (lower < upper) {
    int pivot = partition(arr, lower, upper);
    quick_sort(arr, lower, pivot - 1);
    quick_sort(arr, pivot + 1, upper);
  }
}

void heapify(std::vector<int> &arr, int current_index, int size) {
  int current_max_index = current_index;
  int left_node = current_index * 2 + 1;
  int right_node = current_index * 2 + 2;

  if (left_node < size && arr[left_node] > arr[current_max_index]) {
    current_max_index = left_node;
  }

  if (right_node < size && arr[right_node] > arr[current_max_index]) {
    current_max_index = right_node;
  }

  if (current_max_index != current_index) {
    std::swap(arr[current_index], arr[current_max_index]);
    heapify(arr, current_max_index, size);
  }
}

void heap_sort(std::vector<int> &arr, int size) {
  for (int i = size / 2 - 1; i >= 0; i--) {
    heapify(arr, i, size);
  }

  for (int i = size - 1; i > 0; i--) {
    std::swap(arr[i], arr[0]);
    heapify(arr, 0, i);
  }
}

std::vector<int> get_vector_of_values(std::string &line) {
  std::stringstream ss(line);
  std::string token;
  std::vector<int> arr;
  char delimiter = ',';

  while (std::getline(ss, token, delimiter)) {
    arr.push_back(std::stoi(token));
  }
  return arr;
}

void count_algorithm_execution_time(const std::vector<int> &arr, int size,
                                    const std::string &func_name) {
  std::vector<int> arr_copy = arr;

  auto start = std::chrono::high_resolution_clock::now();

  if (func_name == "quick_sort") {
    quick_sort(arr_copy, 0, size - 1);
  } else if (func_name == "heap_sort") {
    heap_sort(arr_copy, size);
  } else {
    bubble_sort(arr_copy, size);
  }

  auto end = std::chrono::high_resolution_clock::now();
  std::chrono::duration<double, std::milli> elapsed = end - start;

  std::ofstream file(sort_time + func_name + ".csv", std::ios::app);
  file << size << ", " << elapsed.count() << "\n";
}

void perform_test_cases() {
  std::ifstream file(sort_test_cases_file);
  if (file.is_open()) {
    std::string line;
    int size;
    int current_line = 0;

    while (std::getline(file, line)) {
      if (current_line % 2 == 0) {
        size = std::stoi(line);
      } else {
        std::vector<int> arr = get_vector_of_values(line);

        count_algorithm_execution_time(arr, size, "bubble_sort");
        count_algorithm_execution_time(arr, size, "quick_sort");
        count_algorithm_execution_time(arr, size, "heap_sort");
      }
      current_line++;
    }
  }
}

int main() { perform_test_cases(); }
