package slices_basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	sum, err := SumSlice(intSlice)
	assert.NoError(t, err)
	assert.Equal(t, 15, sum)

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	sumFloat, err := SumSlice(floatSlice)
	assert.NoError(t, err)
	assert.Equal(t, 16.5, sumFloat)
}

func TestMinSlice(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	min, err := MinSlice(intSlice)
	assert.NoError(t, err)
	assert.Equal(t, 1, min)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	minFloat, err := MinSlice(floatSlice)
	assert.NoError(t, err)
	assert.Equal(t, 1.1, minFloat)

	emptySlice := []int{}
	_, err = MinSlice(emptySlice)
	assert.Error(t, err)
}

func TestMaxSlice(t *testing.T) {
	intSlice := []int{5, 2, 9, 1, 5, 6}
	max, err := MaxSlice(intSlice)
	assert.NoError(t, err)
	assert.Equal(t, 9, max)

	floatSlice := []float64{3.1, 2.2, 5.5, 1.1, 4.4}
	maxFloat, err := MaxSlice(floatSlice)
	assert.NoError(t, err)
	assert.Equal(t, 5.5, maxFloat)

	emptySlice := []int{}
	_, err = MaxSlice(emptySlice)
	assert.Error(t, err)
}

func TestAvgSlice(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	avg, err := AvgSlice(intSlice)
	assert.NoError(t, err)
	assert.Equal(t, 3, avg)

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	avgFloat, err := AvgSlice(floatSlice)
	assert.NoError(t, err)
	assert.Equal(t, 3.3, avgFloat)

	emptySlice := []int{}
	_, err = AvgSlice(emptySlice)
	assert.Error(t, err)
}
