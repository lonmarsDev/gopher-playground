package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	isDone := make(chan bool, 1)
	isDone <- false
	var wg sync.WaitGroup
	wg.Add(1)
	go robieeFboy(&wg, isDone )
	wg.Add(1)
	go makooFboy(&wg)
	wg.Add(1)
	go daddyJesFboy(&wg)

	wg.Wait()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("\n time elpased : %v \n", elapsed)

}

func robieeFboy(wg *sync.WaitGroup, isdone chan bool) {
	
	for  {
		fmt.Printf("\n Robie fuck girl number # %v\n", i)
		time.Sleep(5 * time.Second)
		select {
			case : isDoneVal:= <- isdone
				if isDoneVal{
					wg.Done()
					return
				}
			default:
					
		}
		time.Sleep(5 * time.Second)
	}
}

func makooFboy(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Printf("\n makoo fuck girl number # %v\n", i)
	}
}

func daddyJesFboy(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Printf("\n DaddyJess fuck girl number # %v\n", i)
	}
}
