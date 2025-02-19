package slices_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	expectedIntSlice := []int{1, 2, 5, 5, 6, 9}
	BubbleSort(intSlice)
	assert.Equal(t, expectedIntSlice, intSlice)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.1, 4.4, 5.5}
	BubbleSort(floatSlice)
	assert.Equal(t, expectedFloatSlice, floatSlice)
}

func TestInsertionSort(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	expectedIntSlice := []int{1, 2, 5, 5, 6, 9}
	InsertionSort(intSlice)
	assert.Equal(t, expectedIntSlice, intSlice)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.1, 4.4, 5.5}
	InsertionSort(floatSlice)
	assert.Equal(t, expectedFloatSlice, floatSlice)
}

func TestSelectionSort(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	expectedIntSlice := []int{1, 2, 5, 5, 6, 9}
	SelectionSort(intSlice)
	assert.Equal(t, expectedIntSlice, intSlice)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.1, 4.4, 5.5}
	SelectionSort(floatSlice)
	assert.Equal(t, expectedFloatSlice, floatSlice)
}

func TestMergeSort(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	expectedIntSlice := []int{1, 2, 5, 5, 6, 9}
	sortedIntSlice := MergeSort(intSlice)
	assert.Equal(t, expectedIntSlice, sortedIntSlice)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	expectedFloatSlice := []float64{1.1, 2.2, 3.1, 4.4, 5.5}
	sortedFloatSlice := MergeSort(floatSlice)
	assert.Equal(t, expectedFloatSlice, sortedFloatSlice)
}
