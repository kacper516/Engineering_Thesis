const fs = require("fs");
const { performance } = require("perf_hooks");

const sort_time =
  "JavaScript/sort_algorithms/sort_execution_time/sort_execution_time_less_jit_darwin/";
const sort_test_cases_file = "/Users/kacper/engineering-thesis/t.csv";

function bubble_sort(arr) {
  for (let i = 0; i < arr.length; i++) {
    let swapped = false;
    for (let j = 0; j < arr.length - i - 1; j++) {
      if (arr[j] > arr[j + 1]) {
        swapped = true;
        [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
      }
    }
    if (!swapped) {
      return; // array sorted
    }
  }
}

function partition(arr, lower, upper) {
  let i = lower - 1;
  let pivot = arr[upper];

  for (let j = lower; j < upper; j++) {
    if (arr[j] < pivot) {
      i++;
      [arr[j], arr[i]] = [arr[i], arr[j]];
    }
  }

  [arr[i + 1], arr[upper]] = [arr[upper], arr[i + 1]];
  return i + 1;
}

function quick_sort(arr, lower, upper) {
  if (lower < upper) {
    let pivot = partition(arr, lower, upper);
    quick_sort(arr, lower, pivot - 1);
    quick_sort(arr, pivot + 1, upper);
  }
}

function heapify(arr, current_index, size) {
  let current_max_index = current_index;
  let left_node = 2 * current_index + 1;
  let right_node = 2 * current_index + 2;

  if (left_node < size && arr[left_node] > arr[current_max_index]) {
    current_max_index = left_node;
  }

  if (right_node < size && arr[right_node] > arr[current_max_index]) {
    current_max_index = right_node;
  }

  if (current_max_index != current_index) {
    [arr[current_max_index], arr[current_index]] = [
      arr[current_index],
      arr[current_max_index],
    ];
    heapify(arr, current_max_index, size);
  }
}

function heap_sort(arr) {
  for (let i = Math.floor(arr.length / 2) - 1; i >= 0; i--) {
    heapify(arr, i, arr.length);
  }

  for (let i = arr.length - 1; i > 0; i--) {
    [arr[0], arr[i]] = [arr[i], arr[0]];
    heapify(arr, 0, i);
  }
}

function get_vector_of_values(line) {
  return line.split(",").map(Number);
}

function count_algorithm_execution_time(arr, size, func_name) {
  let arr_copy = [...arr]; // Create a copy of the array
  let start_time = performance.now();
  if (func_name === "quick_sort") {
    quick_sort(arr_copy, 0, arr_copy.length - 1);
  } else if (func_name === "heap_sort") {
    heap_sort(arr_copy);
  } else {
    bubble_sort(arr_copy);
  }

  let end_time = performance.now();
  let elapsed = end_time - start_time;

  fs.appendFileSync(
    `${sort_time}${func_name}.csv`,
    `${size}, ${elapsed.toFixed(6)}\n`
  );
}

function perform_test_cases() {
  fs.readFile(sort_test_cases_file, "utf8", function (err, data) {
    if (err) {
      console.error("Could not open file:", err);
      return;
    }

    const lines = data.split("\n");
    let size;

    lines.forEach(function (line, index) {
      if (index % 2 === 0) {
        size = parseInt(line, 10);
      } else {
        let arr = get_vector_of_values(line);

        count_algorithm_execution_time(arr, size, "bubble_sort");
        count_algorithm_execution_time(arr, size, "quick_sort");
        count_algorithm_execution_time(arr, size, "heap_sort");
      }
    });
  });
}

perform_test_cases();
