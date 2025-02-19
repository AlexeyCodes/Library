package strings

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
    result, err := Concat("Hello, ", "World!")
    assert.NoError(t, err)
    assert.Equal(t, "Hello, World!", result)

    result, err = Concat("", "World!")
    assert.Error(t, err)
    assert.Equal(t, "", result)
}

func TestContains(t *testing.T) {
    contains, err := Contains("Hello, World!", "World")
    assert.NoError(t, err)
    assert.True(t, contains)

    contains, err = Contains("Hello, World!", "")
    assert.Error(t, err)
    assert.False(t, contains)
}

func TestCount(t *testing.T) {
    count, err := Count("Hello, World!", "o")
    assert.NoError(t, err)
    assert.Equal(t, 2, count)

    count, err = Count("Hello, World!", "")
    assert.Error(t, err)
    assert.Equal(t, 0, count)
}

func TestReverse(t *testing.T) {
    reversed, err := Reverse("Hello, World!")
    assert.NoError(t, err)
    assert.Equal(t, "!dlroW ,olleH", reversed)

    reversed, err = Reverse("")
    assert.Error(t, err)
    assert.Equal(t, "", reversed)
}

func TestToUpper(t *testing.T) {
    upper, err := ToUpper("Hello, World!")
    assert.NoError(t, err)
    assert.Equal(t, "HELLO, WORLD!", upper)

    upper, err = ToUpper("")
    assert.Error(t, err)
    assert.Equal(t, "", upper)
}

func TestToLower(t *testing.T) {
    lower, err := ToLower("Hello, World!")
    assert.NoError(t, err)
    assert.Equal(t, "hello, world!", lower)

    lower, err = ToLower("")
    assert.Error(t, err)
    assert.Equal(t, "", lower)
}

func TestIsPalindrome(t *testing.T) {
    isPalindrome, err := IsPalindrome("madam")
    assert.NoError(t, err)
    assert.True(t, isPalindrome)

    isPalindrome, err = IsPalindrome("hello")
    assert.NoError(t, err)
    assert.False(t, isPalindrome)

    isPalindrome, err = IsPalindrome("")
    assert.Error(t, err)
    assert.False(t, isPalindrome)
}