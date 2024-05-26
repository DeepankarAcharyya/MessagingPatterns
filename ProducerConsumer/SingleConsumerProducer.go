package main

import (
	"fmt"
	"sync"
	"time"
)

func simpleProducer(ch chan string, wg *sync.WaitGroup) {
	for _, c := range []string{
		"Hello Dwight!", "Dammit Dwight", "Oh fucking hell Dwight!",
	} {
		ch <- c
	}
	wg.Done()
}

func simpleConsumer(ch chan string, wg *sync.WaitGroup) {
	for {
		select {
		case msg := <-ch:
			fmt.Printf("Received by the consumer : %s\n", msg)
		case <-time.After(time.Second * 5):
			fmt.Printf("Exiting after waiting for 5 secs")
			wg.Done()
			return
		}
	}

}

func singleConsumerProducer() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	go simpleProducer(ch, &wg)
	wg.Add(1)
	go simpleConsumer(ch, &wg)
	wg.Wait()
	close(ch)

}

func main() {
	singleConsumerProducer()
}
