package numbers

import "errors"

var (
	ErrNegativeFactorial = errors.New("factorial of a negative number is undefined")
	ErrDivisionByZero = errors.New("division by zero")

)