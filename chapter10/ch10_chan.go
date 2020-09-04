package main

import (
	"fmt"
	"time"
)

func pinger(ch chan<- string) {
	for i := 0;; i++ {
		ch <- "ping"
	}
}

func printer(ch <-chan string) {
	for {
		msg := <- ch
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)

	var inp string
	fmt.Scanln(&inp)
}