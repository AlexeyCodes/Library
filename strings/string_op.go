package strings

import (
    "strings"
)

// Concat concatenates two strings.
func Concat(a, b string) (string, error) {
    if a == "" || b == "" {
        return "", ErrEmptyInput // Return an error if either input string is empty
    }
    return a + b, nil
}

// Contains checks if a string contains a substring.
func Contains(s, substr string) (bool, error) {
    if s == "" || substr == "" {
        return false, ErrEmptyInput // Return an error if either input string is empty
    }
    return strings.Contains(s, substr), nil
}

// Count counts the number of occurrences of a substring in a string.
func Count(s, substr string) (int, error) {
    if s == "" || substr == "" {
        return 0, ErrEmptyInput // Return an error if either input string is empty
    }
    return strings.Count(s, substr), nil
}

// Reverse reverses a string.
func Reverse(s string) (string, error) {
    if s == "" {
        return "", ErrEmptyString // Return an error if the input string is empty
    }
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i] // Swap the runes to reverse the string
    }
    return string(runes), nil
}

// ToUpper converts a string to uppercase.
func ToUpper(s string) (string, error) {
    if s == "" {
        return "", ErrEmptyString // Return an error if the input string is empty
    }
    return strings.ToUpper(s), nil
}

// ToLower converts a string to lowercase.
func ToLower(s string) (string, error) {
    if s == "" {
        return "", ErrEmptyString // Return an error if the input string is empty
    }
    return strings.ToLower(s), nil
}

// IsPalindrome checks if a string is a palindrome.
func IsPalindrome(s string) (bool, error) {
    if s == "" {
        return false, ErrEmptyString // Return an error if the input string is empty
    }
    rev, _ := Reverse(s)
    return s == rev, nil
}