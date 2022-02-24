package main

import "fmt"

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		fmt.Println("true")
		return true
	default:
		fmt.Println("false")
		return false
	}
}

func main() {
	for {
		cancelled()
	}
	fmt.Println("exit")
}
