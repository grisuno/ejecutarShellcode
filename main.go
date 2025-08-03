package main

import "fmt"

//go:noinline
func add(a, b int64) int64

func main() {
    sum := add(5, 3)
    fmt.Println("The sum is:", sum)
}