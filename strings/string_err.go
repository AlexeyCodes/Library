package strings

import "errors"

var (
    ErrEmptyString = errors.New("string is empty")
    ErrEmptyInput  = errors.New("one or both strings are empty")
)