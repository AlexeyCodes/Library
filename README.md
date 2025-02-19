Library Project
Version
Current project version: v1.0.0

Description
This project is a set of libraries for working with numbers, strings, and slices in the Go programming language. The project implements various functions for performing mathematical operations, sorting, and working with strings.

Project Structure
main.go: The main file that runs examples of using the libraries.
numbers: Package for working with numbers.
strings: Package for working with strings.
slices_basic: Package for basic slice operations.
slices_sort: Package for sorting slices.
interfaces: Package containing interfaces and type constraints.

Installation
 
Clone the repository:
git clone github.com/AlexeyCodes/Library

Navigate to the project directory
cd Library

Usage
Run the examples by executing the command:
go run main.go

# Package Descriptions

## numbers

The `numbers` package contains functions for performing various mathematical operations:

- `Add`: Adds two numbers.
- `Sub`: Subtracts one number from another.
- `Mult`: Multiplies two numbers.
- `Div`: Divides one number by another.
- `Avg`: Calculates the average of two numbers.
- `Min`: Finds the minimum of two numbers.
- `Max`: Finds the maximum of two numbers.
- `Abs`: Calculates the absolute value of a number.
- `Fact`: Calculates the factorial of a number.
- `IsEven`: Checks if a number is even.
- `IsOdd`: Checks if a number is odd.

## strings

The `strings` package contains functions for working with strings:

- `Concat`: Concatenates two strings.
- `Contains`: Checks if a string contains a substring.
- `Count`: Counts the number of occurrences of a substring in a string.
- `Reverse`: Reverses a string.
- `ToUpper`: Converts a string to uppercase.
- `ToLower`: Converts a string to lowercase.
- `IsPalindrome`: Checks if a string is a palindrome.

## slices_basic

The `slices_basic` package contains functions for basic slice operations:

- `SumSlice`: Calculates the sum of all elements in a slice.
- `MinSlice`: Finds the minimum element in a slice.
- `MaxSlice`: Finds the maximum element in a slice.
- `AvgSlice`: Calculates the average of all elements in a slice.

## slices_sort

The `slices_sort` package contains functions for sorting slices:

- `BubbleSort`: Sorts a slice using the bubble sort algorithm.
- `InsertionSort`: Sorts a slice using the insertion sort algorithm.
- `SelectionSort`: Sorts a slice using the selection sort algorithm.
- `MergeSort`: Sorts a slice using the merge sort algorithm.

## interfaces

The `interfaces` package contains interfaces and type constraints:

- `Number`: A type constraint that allows any numeric type.
- `IntUint`: A type constraint that allows any integer or unsigned integer type.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
