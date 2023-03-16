package main

import (
	"fmt"
	"time"
)

func say(what string) {
	fmt.Println(what)
}

func main() {

	message := "Some message"
	/*
			A goroutine is a function that runs independently of
			the function that started it. Sometimes Go developers explain a
			goroutine as a function that runs as if it were on its own thread.
		    Channel is a pipeline for sending and receiving data.
			Think of it as a socket that runs inside your program.
			Channels provide a way for one goroutine to send structured
			data to another.
	*/
	go say(message)
	fmt.Println("From the main go routine")
	time.Sleep(5 * time.Second)
}
