#include <chrono>
#include <fstream>
#include <iostream>
#include <vector>

const std::string searchTime = "C++/search_algorithms/search_execution_time/";
const std::vector<int> testSizes = {10,     100,     1000,     10000,
                                    100000, 1000000, 10000000, 100000000};

std::vector<int> generate_test_data(int size) {
  std::vector<int> arr(size);
  for (int i = 0; i < size; ++i) {
    arr[i] = i;
  }
  return arr;
}

int linear_search(const std::vector<int> &arr, int targetValue) {
  for (size_t i = 0; i < arr.size(); ++i) {
    if (arr[i] == targetValue) {
      return i;
    }
  }
  return -1;
}

int binary_search(const std::vector<int> &arr, int targetValue) {
  int begin = 0, end = arr.size() - 1;
  while (begin <= end) {
    int mid = begin + (end - begin) / 2;
    if (arr[mid] == targetValue) {
      return mid;
    } else if (arr[mid] < targetValue) {
      begin = mid + 1;
    } else {
      end = mid - 1;
    }
  }
  return -1;
}

void count_algorithm_execution_time(const std::vector<int> &arr, int size,
                                    int targetValue,
                                    const std::string &funcName) {
  auto start = std::chrono::high_resolution_clock::now();

  if (funcName == "binary_search") {
    binary_search(arr, targetValue);
  } else {
    linear_search(arr, targetValue);
  }

  auto end = std::chrono::high_resolution_clock::now();
  std::chrono::duration<double, std::milli> elapsed = end - start;

  std::ofstream file(searchTime + funcName + ".csv", std::ios::app);
  file << size << ", " << elapsed.count() << "\n";
}

void perform_test_cases() {
  for (int size : testSizes) {
    std::vector<int> arr = generate_test_data(size);
    int targetValue = arr.back();

    count_algorithm_execution_time(arr, size, targetValue, "linear_search");
    count_algorithm_execution_time(arr, size, targetValue, "binary_search");
  }
}

int main() {
  perform_test_cases();
  return 0;
}
