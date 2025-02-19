package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result, err := Add(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 3, result)

	resultFloat, err := Add(1.5, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 4.0, resultFloat)
}

func TestSub(t *testing.T) {
	result, err := Sub(5, 3)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)

	resultFloat, err := Sub(5.5, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 3.0, resultFloat)
}

func TestMult(t *testing.T) {
	result, err := Mult(2, 3)
	assert.NoError(t, err)
	assert.Equal(t, 6, result)

	resultFloat, err := Mult(2.5, 2.0)
	assert.NoError(t, err)
	assert.Equal(t, 5.0, resultFloat)
}

func TestDiv(t *testing.T) {
	result, err := Div(6, 3)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)

	resultFloat, err := Div(5.0, 2.0)
	assert.NoError(t, err)
	assert.Equal(t, 2.5, resultFloat)

	_, err = Div(1, 0)
	assert.Error(t, err)
}

func TestMin(t *testing.T) {
	result, err := Min(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)

	resultFloat, err := Min(1.5, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 1.5, resultFloat)
}

func TestMax(t *testing.T) {
	result, err := Max(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, result)

	resultFloat, err := Max(1.5, 2.5)
	assert.NoError(t, err)
	assert.Equal(t, 2.5, resultFloat)
}

func TestAbs(t *testing.T) {
	result, err := Abs(-1)
	assert.NoError(t, err)
	assert.Equal(t, 1, result)

	resultFloat, err := Abs(-1.5)
	assert.NoError(t, err)
	assert.Equal(t, 1.5, resultFloat)
}

func TestFact(t *testing.T) {
	result, err := Fact(5)
	assert.NoError(t, err)
	assert.Equal(t, 120, result)
}

func TestIsEven(t *testing.T) {
	result, err := IsEven(4)
	assert.NoError(t, err)
	assert.True(t, result)

	result, err = IsEven(5)
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestIsOdd(t *testing.T) {
	result, err := IsOdd(4)
	assert.NoError(t, err)
	assert.False(t, result)

	result, err = IsOdd(5)
	assert.NoError(t, err)
	assert.True(t, result)
}
