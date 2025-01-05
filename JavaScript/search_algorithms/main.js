const fs = require("fs");
const { performance } = require("perf_hooks");

const search_time = "JavaScript/search_algorithms/search_execution_time_jit/";
const test_sizes = [10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000];

function generate_test_data(size) {
  let arr = [];
  for (let i = 0; i < size; i++) {
    arr.push(i);
  }
  return [arr, arr[size - 1]];
}

function linear_search(arr, target_value) {
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] === target_value) {
      return i;
    }
  }
  return -1;
}

function binary_search(arr, target_value) {
  let begin = 0;
  let end = arr.length - 1;

  while (begin <= end) {
    let mid = Math.floor(begin + (end - begin) / 2);
    if (arr[mid] === target_value) {
      return mid;
    } else if (arr[mid] < target_value) {
      begin = mid + 1;
    } else {
      end = mid - 1;
    }
  }
  return -1;
}

function count_algorithm_execution_time(
  arr,
  size,
  target_value,
  func_name = "linear_search"
) {
  let start = performance.now();

  switch (func_name) {
    case "binary_search":
      binary_search(arr, target_value);
      break;
    default:
      linear_search(arr, target_value);
  }

  let end = performance.now();
  let elapsed = end - start;

  fs.appendFileSync(
    `${search_time}${func_name}.csv`,
    `${size}, ${elapsed.toFixed(6)}\n`
  );
}

function perform_test_cases() {
  for (let i = 0; i < test_sizes.length; i++) {
    let size = test_sizes[i];
    let [arr, target_value] = generate_test_data(size);
    count_algorithm_execution_time(arr, size, target_value, "linear_search");
    count_algorithm_execution_time(arr, size, target_value, "binary_search");
  }
}

perform_test_cases();
