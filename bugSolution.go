```go
func main() {
    var i int
    fmt.Println(i) // Output: 0

    i = 10
    fmt.Println(i) // Output: 10

    var finalI int
    defer func() {
        finalI = i + 1 // Modify a copy of i
        fmt.Println("Deferred finalI:", finalI) // Output: Deferred finalI: 11
    }()

    fmt.Println("Final i:", i) // Output: Final i: 10
}
```