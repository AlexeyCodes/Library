package examples

import (
	"fmt"
	"Library/numbers"
	"Library/slices_basic"
	"Library/slices_sort"
	"Library/strings"
)

func RunExamples() {
    // Examples of functions from the numbers package
    fmt.Println("Numbers package examples:")
    sum, err := numbers.Add(1, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sum:", sum)
    }

    sub, err := numbers.Sub(5, 3)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Subtraction:", sub)
    }

    mult, err := numbers.Mult(2, 3)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Multiplication:", mult)
    }

    div, err := numbers.Div(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division:", div)
    }

    avg, err := numbers.Avg(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Average:", avg)
    }

    min, err := numbers.Min(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Minimum:", min)
    }

    max, err := numbers.Max(4, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Maximum:", max)
    }

    abs, err := numbers.Abs(-4)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Absolute value:", abs)
    }

    fact, err := numbers.Fact(5)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Factorial:", fact)
    }

    isEven, err := numbers.IsEven(4)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Is even:", isEven)
    }

    isOdd, err := numbers.IsOdd(3)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Is odd:", isOdd)
    }

    // Examples of functions from the strings package
    fmt.Println("\nStrings package examples:")
    concat, err := strings.Concat("Hello, ", "World!")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Concatenated string:", concat)
    }

    contains, err := strings.Contains("Hello, World!", "World")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Contains substring:", contains)
    }

    count, err := strings.Count("Hello, World!", "o")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Count of substring:", count)
    }

    reversed, err := strings.Reverse("Hello, World!")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Reversed string:", reversed)
    }

    upper, err := strings.ToUpper("Hello, World!")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Uppercase string:", upper)
    }

    lower, err := strings.ToLower("Hello, World!")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Lowercase string:", lower)
    }

    isPalindrome, err := strings.IsPalindrome("madam")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Is palindrome:", isPalindrome)
    }

    isPalindrome, err = strings.IsPalindrome("hello")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Is palindrome:", isPalindrome)
    }

    // Examples of functions from the slices package
    fmt.Println("\nSlices package examples:")
    intSlice := []int{5, 2, 9, 1, 5, 6}
    fmt.Println("Original int slice:", intSlice)

    slices_sort.BubbleSort(intSlice)
    fmt.Println("Bubble sorted int slice:", intSlice)

    intSlice = []int{5, 2, 9, 1, 5, 6}
    slices_sort.InsertionSort(intSlice)
    fmt.Println("Insertion sorted int slice:", intSlice)

    intSlice = []int{5, 2, 9, 1, 5, 6}
    slices_sort.SelectionSort(intSlice)
    fmt.Println("Selection sorted int slice:", intSlice)

    intSlice = []int{5, 2, 9, 1, 5, 6}
    sortedIntSlice := slices_sort.MergeSort(intSlice)
    fmt.Println("Merge sorted int slice:", sortedIntSlice)

    sumSlice, err := slices_basic.SumSlice(intSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sum of int slice:", sumSlice)
    }

    minSlice, err := slices_basic.MinSlice(intSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Min of int slice:", minSlice)
    }

    maxSlice, err := slices_basic.MaxSlice(intSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Max of int slice:", maxSlice)
    }

    avgSlice, err := slices_basic.AvgSlice(intSlice)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Avg of int slice:", avgSlice)
    }
}