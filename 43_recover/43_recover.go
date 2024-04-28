package main

import "fmt"

func mayPanic() {
    panic("a problem")
}

func main() {

    defer func() {
        if r := recover(); r != nil {

            fmt.Println("Recovered. Error:\n", r)
        }
    }()
	fmt.Println("It goes here...")
    mayPanic()

    fmt.Println("After mayPanic()")
}