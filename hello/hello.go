package main

import (
    "fmt"

    "example.com/greeting"
)

func main() {
    message := greeting.Greet("Sam")
    fmt.Println(message)
}
