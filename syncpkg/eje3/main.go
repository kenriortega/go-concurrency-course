package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	wg.Add(1)
	go func() {
		defer wg.Done()

		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		cond.L.Unlock()
	}()
	cond.L.Lock()

	sharedRsc["rsc1"] = "foo"
	cond.Signal()
	cond.L.Unlock()
	wg.Wait()
}
