package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var s sync.WaitGroup

func handleMessage(c chan int, index int) {
	select {
	case y := <-c:
		fmt.Printf("index %d: %d\n", index, y)
	}

	s.Done()
}

func main() {

	c := make(chan int)
	for i := 0; i < 20; i++ {
		s.Add(1)
		go handleMessage(c, i)
	}

	for i := 0; i < 20; i++ {
		c <- rand.Int() % 1000
	}

	s.Wait()
}
