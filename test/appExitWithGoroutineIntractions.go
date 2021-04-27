package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		// IT will fail to run properly and will give dealock error if wg is passed as value and not referance
		// (Shown below)
		// go doit(i, done, wg)
		go doit(i, done, &wg)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doit(workerId int, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerId)
}
