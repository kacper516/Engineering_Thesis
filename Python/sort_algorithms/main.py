import time

sort_time = "Python/sort_algorithms/sort_execution_time_darwin/"
sort_test_cases_file = "/Users/kacper/engineering-thesis/t.csv"


def bubble_sort(arr):
    for i in range(len(arr)):
        swapped = False
        for j in range(0, len(arr) - i - 1):
            if arr[j] > arr[j + 1]:
                swapped = True
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
        if not swapped:
            return  # there's sorted arr


def partition(arr, lower, upper):
    pivot = arr[upper]
    i = lower - 1

    for j in range(lower, upper):
        if arr[j] <= pivot:
            i += 1
            arr[j], arr[i] = arr[i], arr[j]

    arr[i + 1], arr[upper] = arr[upper], arr[i + 1]
    return i + 1


def quick_sort(arr, lower, upper):
    if lower < upper:
        pivot = partition(arr, lower, upper)

        quick_sort(arr, lower, pivot - 1)
        quick_sort(arr, pivot + 1, upper)


def heapify(arr, current_index, size):
    current_max_index = current_index
    left_node = current_index * 2 + 1
    right_index = current_index * 2 + 2

    if left_node < size and arr[current_max_index] < arr[left_node]:
        current_max_index = left_node

    if right_index < size and arr[current_max_index] < arr[right_index]:
        current_max_index = right_index

    if current_max_index != current_index:
        arr[current_index], arr[current_max_index] = (
            arr[current_max_index],
            arr[current_index],
        )
        heapify(arr, current_max_index, size)


def heap_sort(arr):
    size = len(arr)

    for i in range(size // 2 - 1, -1, -1):
        heapify(arr, i, size)

    for i in range(size - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapify(arr, 0, i)


def count_algorithm_execution_time(arr, size, func_name="bubble_sort"):
    time_start = time.perf_counter()
    if func_name == "bubble_sort":
        bubble_sort(arr)
    elif func_name == "quick_sort":
        quick_sort(arr, 0, size - 1)
    elif func_name == "heap_sort":
        heap_sort(arr)

    time_end = time.perf_counter()
    with open(f"{sort_time}/{func_name}.csv", "a") as file:
        file.write(f"{size}, {(time_end - time_start)*1000:.6f}\n")


def perform_test_cases(lines):
    for current_line_idx in range(0, len(lines), 2):
        size = int(lines[current_line_idx].strip())
        arr = [int(number) for number in lines[current_line_idx + 1].split(",")]

        count_algorithm_execution_time(arr.copy(), size)
        count_algorithm_execution_time(arr.copy(), size, "quick_sort")
        count_algorithm_execution_time(arr.copy(), size, "heap_sort")


if __name__ == "__main__":
    with open(file=sort_test_cases_file, mode="r") as file:
        lines = file.readlines()

        perform_test_cases(lines)
