use std::fs::OpenOptions;
use std::io::Write;
use std::time::Instant;

const SEARCH_TIME: &str = "search_execution_time/";
const TEST_SIZES: [usize; 8] = [10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000];

fn generate_test_data(size: usize) -> (Vec<usize>, usize) {
    let arr: Vec<usize> = (0..size).collect();
    (arr.clone(), arr[size - 1])
}

fn linear_search(arr: &[usize], target_value: usize) -> isize {
    for (idx, &value) in arr.iter().enumerate() {
        if value == target_value {
            return idx as isize;
        }
    }
    -1
}

fn binary_search(arr: &[usize], target_value: usize) -> isize {
    let (mut begin, mut end) = (0, arr.len() - 1);
    while begin <= end {
        let mid = begin + (end - begin) / 2;
        if arr[mid] == target_value {
            return mid as isize;
        } else if arr[mid] < target_value {
            begin = mid + 1;
        } else {
            end = mid - 1;
        }
    }
    -1
}

fn count_algorithm_execution_time(
    arr: &[usize],
    size: usize,
    target_value: usize,
    func_name: &str,
) {
    let start = Instant::now();

    match func_name {
        "binary_search" => {
            binary_search(arr, target_value);
        }
        _ => {
            linear_search(arr, target_value);
        }
    }

    let elapsed = start.elapsed().as_secs_f64() * 1000.0;

    let mut file = OpenOptions::new()
        .append(true)
        .create(true)
        .open(format!("{}{}.csv", SEARCH_TIME, func_name))
        .expect("Cannot open file");

    writeln!(file, "{}, {:.6}", size, elapsed).expect("Cannot write to file");
}

fn perform_test_cases() {
    for &size in &TEST_SIZES {
        let (arr, target_value) = generate_test_data(size);
        count_algorithm_execution_time(&arr, size, target_value, "linear_search");
        count_algorithm_execution_time(&arr, size, target_value, "binary_search");
    }
}

fn main() {
    perform_test_cases();
}
