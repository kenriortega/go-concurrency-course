package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}
func main() {
	fun("Direct call")

	// goroutine func call
	go fun("goroutine-1")
	// anonymous func
	go func() {
		fun("goroutine-2")
	}()
	// goroutine with func value call
	fv := fun
	go fv("goroutine-3")
	// wait for a goroutine to end
	fmt.Println("wait for goroutines...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done ...")
}
