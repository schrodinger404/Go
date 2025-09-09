package main

import (
	"fmt"
	"sync"
	//"time"
)


func tasks(i int, w *sync.WaitGroup)  {
	// waitGroup Done() function
	defer w.Done()
	fmt.Println("task:", i)
	
}

func main() {

	// for i := 0; i <= 10; i++ {

	// 	//anonimus func, if we add go infront of func it'll run concurrent
	// 	go func(i int) {
	// 		fmt.Println("task:", i)
	// 	}(i)

	// }

	// // since main also runs fast, just to sleep it for 2 seconds
	// time.Sleep(time.Second * 2)

	// this waitGroup has three functions that need to be defined
	
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {

		//waitGroup Add() function
		wg.Add(1)
		go tasks(i, &wg)
		
	}

	// waitGroup Wait() function
	wg.Wait()
}