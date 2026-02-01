package main

import "fmt"

func main() {
  // Anonymous function assigned to variable
  greet := func(name string) {
    fmt.Println("Hello,", name)
  }

  // Call the anonymous function
  greet("Alice")
  greet("Bob")

  // Anonymous function with return value
  add := func(a, b int) int {
    return a + b
  }

  result := add(10, 20)
  fmt.Println("10 + 20 =", result)

  // Passing anonymous function as argument
  numbers := []int{1, 2, 3, 4, 5}

  apply := func(nums []int, fn func(int) int) []int {
    result := make([]int, len(nums))
    for i, n := range nums {
      result[i] = fn(n)
    }
    return result
  }

  doubled := apply(numbers, func(n int) int {
    return n * 2
  })

  fmt.Println("Original:", numbers)
  fmt.Println("Doubled:", doubled)
}
