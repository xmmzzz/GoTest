package main

import (
	"fmt"
	"time"
)

func say(s string,  c chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
	c <- true
	return
}

func main() {
	c := make(chan bool, 1)
	go say("hello", c)

	say("world", c)
	<-c
	<-c
}
