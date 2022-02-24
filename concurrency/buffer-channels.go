package main

import (
	"errors"
	"fmt"
)

func main() {
	ch := make(chan int, 100)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println(errors.New("a") == errors.New("a"))
}
