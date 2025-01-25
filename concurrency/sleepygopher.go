package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("... ", id, " snore...")
	// sends the id on a channel named c
	c <- id
}

func SleepMain() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		// recieves the channel's value
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
	time.Sleep(4 * time.Second)
}
