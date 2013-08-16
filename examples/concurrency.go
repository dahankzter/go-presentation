package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	now := time.Now()

	go func() {
		fmt.Printf("Catching up after %v\n", time.Since(now))
		done <- true
	}()

	fmt.Printf("Moving along after %v\n", time.Since(now))

	<-done
}
