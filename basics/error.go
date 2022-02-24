package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(errors.New("a") == errors.New("a"))
}
