import time

search_time = "Python/search_algorithms/search_execution_time_macos/"

# number of elements inside array
test_sizes = [10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000]


def generate_test_data(size):
    """Generates array with properly size, and target for the last element inside it"""
    arr = list(range(size))
    target_value = arr[-1]
    return arr, target_value


def linear_search(arr, target_value):
    for idx in range(len(arr)):
        if arr[idx] == target_value:
            return idx
    return -1


def binary_search(arr, target_value):
    begin = 0
    end = len(arr) - 1

    while begin <= end:
        mid = begin + (end - begin) // 2
        if arr[mid] == target_value:
            return mid
        elif arr[mid] < target_value:
            begin = mid + 1
        else:
            end = mid - 1
    return -1


def count_algorithm_execution_time(arr, size, target_value, func_name="linear_search"):
    func = linear_search  # by default

    if func_name == "binary_search":
        func = binary_search

    time_start = time.perf_counter()
    func(arr, target_value)
    time_end = time.perf_counter()
    with open(f"{search_time}/{func_name}.csv", "a") as file:
        file.write(f"{size}, {(time_end - time_start)*1000:.6f}\n")


def perform_test_cases():
    for size in test_sizes:
        arr, target_value = generate_test_data(size)

        count_algorithm_execution_time(arr, size, target_value)
        count_algorithm_execution_time(arr, size, target_value, "binary_search")


if __name__ == "__main__":
    perform_test_cases()
