package main

import (
	"fmt"
	"time"
)

func say(s string, duration time.Duration, c chan bool) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, ":", s)
		time.Sleep(duration)
	}
	c <- true
	return
}

func main() {
	c := make(chan bool, 2)
	go say("hello", 50*time.Millisecond, c)

	say("world", 100*time.Millisecond, c)
	<-c
	<-c
}
