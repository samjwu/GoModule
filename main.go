package main

import (
    "fmt"
    "log"

    "example.com/greeting"
)

func main() {
    log.SetPrefix("greeting: ")

    // https://pkg.go.dev/log#Flags
    log.SetFlags(0)

    message, err := greeting.Greet("")

    if err != nil {
        // https://pkg.go.dev/log#Fatal
        log.Fatal(err)
    }

    fmt.Println(message)
}
