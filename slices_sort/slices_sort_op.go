package slices_sort

import "Library/interfaces"

// BubbleSort sorts a slice using the bubble sort algorithm.
func BubbleSort[T interfaces.Number](slice []T) {
    n := len(slice)

    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if slice[j] > slice[j+1] {
                slice[j], slice[j+1] = slice[j+1], slice[j] // Swap elements if they are in the wrong order
            }
        }
    }
}

// InsertionSort sorts a slice using the insertion sort algorithm.
func InsertionSort[T interfaces.Number](slice []T) {
    n := len(slice)

    for i := 1; i < n; i++ {
        key := slice[i]
        j := i - 1

        for j >= 0 && slice[j] > key {
            slice[j+1] = slice[j] // Shift elements to the right to make room for the key
            j--
        }

        slice[j+1] = key // Insert the key in its correct position
    }
}

// SelectionSort sorts a slice using the selection sort algorithm.
func SelectionSort[T interfaces.Number](slice []T) {
    n := len(slice)

    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if slice[j] < slice[minIdx] {
                minIdx = j // Find the index of the minimum element
            }
        }

        slice[i], slice[minIdx] = slice[minIdx], slice[i] // Swap the minimum element with the first element
    }
}

// MergeSort sorts a slice using the merge sort algorithm.
func MergeSort[T interfaces.Number](slice []T) []T {
    if len(slice) < 2 {
        return slice // Base case: a slice with less than 2 elements is already sorted
    }

    mid := len(slice) / 2
    left := MergeSort(slice[:mid]) // Recursively sort the left half
    right := MergeSort(slice[mid:]) // Recursively sort the right half

    return merge(left, right) // Merge the sorted halves
}

// merge merges two sorted slices into a single sorted slice.
func merge[T interfaces.Number](left, right []T) []T {
    result := make([]T, 0, len(left)+len(right))

    i, j := 0, 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i]) // Append the smaller element to the result
            i++
        } else {
            result = append(result, right[j]) // Append the smaller element to the result
            j++
        }
    }

    result = append(result, left[i:]...) // Append any remaining elements from the left slice
    result = append(result, right[j:]...) // Append any remaining elements from the right slice
    return result
}