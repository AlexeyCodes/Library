package numbers

import (
    "Library/interfaces"
)

// Add adds two numbers.
func Add[T interfaces.Number](a, b T) (T, error) {
    return a + b, nil
}

// Sub subtracts the second number from the first.
func Sub[T interfaces.Number](a, b T) (T, error) {
    return a - b, nil
}

// Mult multiplies two numbers.
func Mult[T interfaces.Number](a, b T) (T, error) {
    return a * b, nil
}

// Div divides the first number by the second.
func Div[T interfaces.Number](a, b T) (T, error) {
    if b == 0 {
        return 0, ErrDivisionByZero
    }
    return a / b, nil
}

// Avg calculates the average of two numbers.
func Avg[T interfaces.Number](a, b T) (T, error) {
    return (a + b) / 2, nil
}

// Min returns the minimum of two numbers.
func Min[T interfaces.Number](a, b T) (T, error) {
    if a < b {
        return a, nil
    }
    return b, nil
}

// Max returns the maximum of two numbers.
func Max[T interfaces.Number](a, b T) (T, error) {
    if a > b {
        return a, nil
    }
    return b, nil
}

// Abs returns the absolute value of a number.
func Abs[T interfaces.Number](a T) (T, error) {
    if a < 0 {
        return -a, nil
    }
    return a, nil
}

// Fact calculates the factorial of a number.
func Fact[T interfaces.Number](a T) (T, error) {
    if a < 0 {
        return 0, ErrNegativeFactorial
    }
    var result T = 1
    for i := T(1); i <= a; i++ {
        result *= i
    }
    return result, nil
}

// IsEven checks if a number is even.
func IsEven[T interfaces.IntUint](a T) (bool, error) {
    return a%2 == 0, nil
}

// IsOdd checks if a number is odd.
func IsOdd[T interfaces.IntUint](a T) (bool, error) {
    return a%2 != 0, nil
}