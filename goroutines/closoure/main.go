package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// incr := func(wg *sync.WaitGroup) {
	// 	var i int
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		i++
	// 		fmt.Println("value i: ", i)
	// 	}()
	// 	fmt.Println("return from func")
	// }
	// incr(&wg)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("i: ", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Done")
}
