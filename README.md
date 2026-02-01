# Go Lesson 13: Anonymous Functions & Closures

Demo code from Go Tutorial for Beginners - Lesson 13.

## Files

- `anon_basics.go` - Anonymous functions assigned to variables, passing as arguments
- `anon_iife.go` - Immediately Invoked Function Expressions (IIFE)
- `closure_basics.go` - Closures capturing outer scope variables
- `closure_factory.go` - Factory pattern with closures

## Running the Examples

```bash
# Anonymous function basics
go run anon_basics.go

# IIFE examples
go run anon_iife.go

# Closure basics
go run closure_basics.go

# Factory pattern
go run closure_factory.go
```

## Key Concepts

### Anonymous Functions
```go
greet := func(name string) {
    fmt.Println("Hello,", name)
}
greet("Alice")
```

### IIFE (Immediately Invoked Function Expression)
```go
func() {
    fmt.Println("I run immediately!")
}()
```

### Closures
```go
count := 0
increment := func() int {
    count++
    return count
}
fmt.Println(increment()) // 1
fmt.Println(increment()) // 2
```

### Factory Pattern
```go
func multiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}
double := multiplier(2)
fmt.Println(double(5)) // 10
```

## Watch the Tutorial

[Go Tutorial for Beginners #13 - Anonymous Functions & Closures](https://youtube.com/@CelesteAI)

## License

MIT
