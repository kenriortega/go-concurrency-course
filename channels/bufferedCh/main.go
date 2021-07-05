package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
