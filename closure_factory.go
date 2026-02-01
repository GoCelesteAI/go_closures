package main

import "fmt"

// Factory function - returns a function
func multiplier(factor int) func(int) int {
  return func(n int) int {
    return n * factor
  }
}

// Adder factory
func adder(amount int) func(int) int {
  return func(n int) int {
    return n + amount
  }
}

// Practical example: logger with prefix
func makeLogger(prefix string) func(string) {
  return func(message string) {
    fmt.Printf("[%s] %s\n", prefix, message)
  }
}

func main() {
  double := multiplier(2)
  triple := multiplier(3)
  times10 := multiplier(10)

  fmt.Println("=== Multiplier Factory ===")
  fmt.Println("double(5) =", double(5))
  fmt.Println("triple(5) =", triple(5))
  fmt.Println("times10(5) =", times10(5))

  fmt.Println("\n=== Adder Factory ===")
  add5 := adder(5)
  add100 := adder(100)

  fmt.Println("add5(10) =", add5(10))
  fmt.Println("add100(10) =", add100(10))

  fmt.Println("\n=== Logger Factory ===")
  infoLog := makeLogger("INFO")
  errorLog := makeLogger("ERROR")
  debugLog := makeLogger("DEBUG")

  infoLog("Server started")
  errorLog("Connection failed")
  debugLog("Variable x = 42")

  fmt.Println("\n=== Chaining ===")
  value := 5
  value = double(value)
  value = triple(value)
  value = add5(value)
  fmt.Println("5 -> double -> triple -> add5 =", value)
}
