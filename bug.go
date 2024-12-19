```go
func main() {
    var i int
    fmt.Println(i) // Output: 0

    i = 10
    fmt.Println(i) // Output: 10

    defer func() {
        i++
        fmt.Println("Deferred i:", i) // Output: Deferred i: 11
    }()

    fmt.Println("Final i:", i) // Output: Final i: 10
}
```