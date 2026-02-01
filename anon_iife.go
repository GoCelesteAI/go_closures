package main

import "fmt"

func main() {
  // IIFE - Immediately Invoked Function Expression
  func() {
    fmt.Println("I run immediately!")
  }()

  // IIFE with parameters
  func(name string) {
    fmt.Println("Hello,", name)
  }("World")

  // IIFE with return value
  result := func(a, b int) int {
    return a + b
  }(10, 20)

  fmt.Println("Result:", result)

  // Practical use: initialize complex value
  config := func() map[string]string {
    m := make(map[string]string)
    m["host"] = "localhost"
    m["port"] = "8080"
    m["debug"] = "true"
    return m
  }()

  fmt.Println("\nConfig initialized:")
  for k, v := range config {
    fmt.Printf("  %s: %s\n", k, v)
  }

  // IIFE for scope isolation
  x := 10
  func() {
    x := 100
    fmt.Println("\nInside IIFE, x =", x)
  }()
  fmt.Println("Outside IIFE, x =", x)
}
