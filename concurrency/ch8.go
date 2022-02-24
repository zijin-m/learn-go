package main

import "fmt"

var balances = make(chan int) // receive balance

var e = make(chan struct{})

func main() {
	go teller()

	<-e

}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case balances <- balance:
			fmt.Println("read balacne")
		}
	}
}
