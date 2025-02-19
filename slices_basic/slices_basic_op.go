package slices_basic

import (
	"Library/interfaces"
	"errors"
)


// SumSlice calculates the sum of all elements in a slice.
func SumSlice[T interfaces.Number](slice []T) (T, error) {
	var sum T

	for _, v := range slice {
		sum += v // Add each element to the sum
	}

	return sum, nil

}


// MinSlice returns the minimum element of the slice 
func MinSlice[T interfaces.Number](slice []T) (T, error) {

	var min T

	if len(slice) == 0 {
		return 0, ErrEmptySlice // Return an error if the slice is empty
	}

	min = slice[0]

	for _, v := range slice {
		if v < min {
			min = v // Update min if a smaller element is found
		}
	}

	return min, nil

}


// MaxSlice returns the maximum element in a slice.
func MaxSlice[T interfaces.Number](slice []T) (T, error) {
	var max T

	if len(slice) == 0 {
		return 0, ErrEmptySlice // Return an error if the slice is empty
	}

	max = slice[0]

	for _, v := range slice {
		if v > max {
			max = v // Update max if a larger element is found
		}
	}

	return max, nil

}


// AvgSlice calculates the average of all elements in a slice.
func AvgSlice[T interfaces.Number](slice []T) (T, error) {

	if len(slice) == 0 {
		return 0, errors.New("slice is empty") // Return an error if the slice is empty
	}

	sum, _ := SumSlice(slice)  // Calculate the sum of the slice

	return sum / T(len(slice)), nil // Return the average

}
