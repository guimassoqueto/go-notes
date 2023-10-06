package main

import (
	"fmt"
	"sync"
)


func main() {
	counter := 0
	processes := 100

	var wg sync.WaitGroup
	wg.Add(processes)
	for i := 0; i < processes; i++ {
		go func(){
			v := counter
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}