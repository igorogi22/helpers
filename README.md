# Helpers

A comprehensive collection of helper functions for common operations in Go. This module is organized into different categories to make it easy to find and use the utilities you need.

---

## Installation

```bash
go get igorogi22/helpers
```

## Usage

```go
import helpers "igorogi22/helpers"
```

---

## Arrays

A collection of functions for array and slice operations in Go, including numeric array generation, random element selection, array reversal, and shuffling operations.

### Functions

#### `CreateNumericArray`

Generates a sequential numeric array starting from a specified value.

**Signature:**
```go
func CreateNumericArray(size int, start ...int) []int
```

**Example:**

```go
import helpers "igorogi22/helpers"

// Generate array from 1 to 5
arr1 := helpers.CreateNumericArray(5)
// Result: [1, 2, 3, 4, 5]

// Generate array starting from 10
arr2 := helpers.CreateNumericArray(3, 10)
// Result: [10, 11, 12]

// Generate array starting from -5
arr3 := helpers.CreateNumericArray(4, -5)
// Result: [-5, -4, -3, -2]
```

**Use Cases:**
- Creating test data for numerical ranges
- Generating sequential IDs or indices
- Building numbered lists programmatically

---

#### `GetRandomItem`

Selects a random element from the provided array or slice. Works with any slice type.

**Signature:**
```go
func GetRandomItem[T any](arr []T) T
```

**Example:**

```go
import helpers "igorogi22/helpers"

// Get random string from slice
fruits := []string{"apple", "banana", "orange"}
fruit := helpers.GetRandomItem(fruits)
// Result: randomly returns "apple", "banana", or "orange"

// Get random number from slice
numbers := []int{10, 20, 30, 40}
num := helpers.GetRandomItem(numbers)
// Result: randomly returns one of the numbers
```

**Use Cases:**
- Random item selection from a list
- Picking random samples for testing
- Implementing randomized algorithms

---

#### `Reverse`

Reverses the elements of an array in place. This is a generic function that works with any slice type.

**Signature:**
```go
func Reverse[T any](arr []T) []T
```

**Example:**

```go
import helpers "igorogi22/helpers"

// Reverse a slice of integers
numbers := []int{1, 2, 3, 4, 5}
helpers.Reverse(numbers)
// numbers is now: [5, 4, 3, 2, 1]

// Reverse strings in place
words := []string{"hello", "world", "go"}
helpers.Reverse(words)
// words is now: ["go", "world", "hello"]
```

**Notes:**
- This function modifies the original slice
- The operation has O(n) time complexity
- Empty or single-element slices remain unchanged

**Use Cases:**
- Reversing arrays before processing
- Implementing palindrome checks
- Converting between little-endian and big-endian formats

---

#### `ShuffleArray`

Randomly shuffles the elements of an array using the Fisher-Yates shuffle algorithm. This is a generic function that works with any slice type.

**Signature:**
```go
func ShuffleArray[T any](arr []T) []T
```

**Example:**

```go
import helpers "igorogi22/helpers"

// Shuffle an array of numbers
numbers := []int{1, 2, 3, 4, 5}
helpers.ShuffleArray(numbers)
// numbers might now be: [3, 1, 5, 2, 4] (or any other permutation)

// Shuffle strings
fruits := []string{"apple", "banana", "orange", "grape"}
helpers.ShuffleArray(fruits)
// fruits is now randomly ordered

// Shuffle a larger array
deck := helpers.CreateNumericArray(52)  // Simulate card deck
helpers.ShuffleArray(deck)
// deck is now randomly ordered
```

**Algorithm:**
- Uses Fisher-Yates shuffle for uniform random distribution
- Each element has equal probability of ending up in any position
- Time complexity: O(n) where n is the array length

**Use Cases:**
- Shuffling card decks for games
- Randomizing question order in quizzes
- Creating random permutations for algorithms
- Data anonymization by randomizing order

---

## Coming Soon

More helper categories will be added in future versions:
- Strings
- Numbers
- Date/Time
- Collections
- Validation
