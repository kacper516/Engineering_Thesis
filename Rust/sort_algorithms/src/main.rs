use std::{
    fs::{self, OpenOptions},
    io::Write,
    time::Instant,
};

const SORT_TIME: &str = "sort_execution_time/";
const SORT_TEST_CASES_FILE: &str =
    "C:/Users/Kacper/Documents/Engeenering_thesis/sort_test_cases.csv";

fn bubble_sort(arr: &mut Vec<i32>, size: usize) {
    for i in 0..size {
        let mut swapped = false;
        for j in 0..size - i - 1 {
            if arr[j] > arr[j + 1] {
                arr.swap(j, j + 1);
                swapped = true;
            }
        }
        if !swapped {
            break;
        }
    }
}

fn partition(arr: &mut Vec<i32>, lower: isize, upper: isize) -> isize {
    let pivot = arr[upper as usize];
    let mut i = lower - 1;

    for j in lower..upper {
        if arr[j as usize] < pivot {
            i += 1;
            arr.swap(i as usize, j as usize);
        }
    }

    arr.swap((i + 1) as usize, upper as usize);
    i + 1
}

fn quick_sort(arr: &mut Vec<i32>, lower: isize, upper: isize) {
    if lower < upper {
        let pivot = partition(arr, lower, upper);
        quick_sort(arr, lower, pivot - 1);
        quick_sort(arr, pivot + 1, upper);
    }
}

fn heapify(arr: &mut Vec<i32>, size: usize, current_index: usize) {
    let mut largest = current_index;
    let left = 2 * current_index + 1;
    let right = 2 * current_index + 2;

    if left < size && arr[left] > arr[largest] {
        largest = left;
    }

    if right < size && arr[right] > arr[largest] {
        largest = right;
    }

    if largest != current_index {
        arr.swap(current_index, largest);
        heapify(arr, size, largest);
    }
}

fn heap_sort(arr: &mut Vec<i32>, size: usize) {
    for i in (0..size / 2 - 1).rev() {
        heapify(arr, size, i);
    }

    for i in (1..size - 1).rev() {
        arr.swap(0, i);
        heapify(arr, i, 0);
    }
}

fn count_algorithm_execution_time(arr: &mut Vec<i32>, size: usize, func_name: &str) {
    let start = Instant::now();
    match func_name {
        "quick_sort" => {
            quick_sort(arr, 0, (size - 1) as isize);
        }
        "heap_sort" => {
            heap_sort(arr, size);
        }
        _ => {
            bubble_sort(arr, size);
        }
    }
    let elapsed = start.elapsed().as_secs_f64() * 1000.0;

    let mut file = OpenOptions::new()
        .append(true)
        .create(true)
        .open(format!("{}{}.csv", SORT_TIME, func_name))
        .expect("Cannot open file");

    writeln!(file, "{}, {:.6}", size, elapsed).expect("Cannot write to file");
}

fn perform_test_cases() {
    let file_contents =
        fs::read_to_string(SORT_TEST_CASES_FILE).expect("Error while reading to string from file.");
    let lines: Vec<&str> = file_contents.lines().collect();

    for i in (0..lines.len()).step_by(2) {
        let size = lines[i].parse::<usize>().unwrap();
        let arr: Vec<i32> = lines[i + 1]
            .split(',')
            .map(|x| x.parse::<i32>().unwrap())
            .collect();

        let mut arr_bubble = arr.clone();
        count_algorithm_execution_time(&mut arr_bubble, size, "bubble_sort");

        let mut arr_quick = arr.clone();
        count_algorithm_execution_time(&mut arr_quick, size, "quick_sort");

        let mut arr_heap = arr.clone();
        count_algorithm_execution_time(&mut arr_heap, size, "heap_sort");
    }
}

fn main() {
    perform_test_cases();
}
