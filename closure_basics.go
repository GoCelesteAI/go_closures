package main

import "fmt"

func main() {
  // A closure captures variables from its outer scope
  message := "Hello from closure"

  printMessage := func() {
    fmt.Println(message)
  }

  printMessage()

  // The closure sees changes to captured variables
  message = "Message changed!"
  printMessage()

  fmt.Println("\n=== Counter Closure ===")
  count := 0

  increment := func() int {
    count++
    return count
  }

  fmt.Println("Count:", increment())
  fmt.Println("Count:", increment())
  fmt.Println("Count:", increment())

  fmt.Println("\n=== Multiple Closures ===")

  makeCounter := func() func() int {
    n := 0
    return func() int {
      n++
      return n
    }
  }

  counter1 := makeCounter()
  counter2 := makeCounter()

  fmt.Println("Counter1:", counter1())
  fmt.Println("Counter1:", counter1())
  fmt.Println("Counter2:", counter2())
  fmt.Println("Counter1:", counter1())
}
