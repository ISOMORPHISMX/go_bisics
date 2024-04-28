package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()
	go func() { messages <-"ping again" }()

    msg := <- messages
    fmt.Println(msg)

	msg =  <- messages
	fmt.Println(msg)
}