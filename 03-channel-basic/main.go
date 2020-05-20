package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Running Unbuffer Channel.")
	unbuffered()
	fmt.Println("Running buffer Channel.")
	buffered()
	fmt.Println("Running Uni Directional Channel.")
	uni_directional()
	fmt.Println("Running Closed Channel")
	close_channel()
	fmt.Println("Running If Channel")
	channel_if()
	fmt.Println("Running For Channel")
	channel_for()

}

func unbuffered() {
	wg := &sync.WaitGroup{}
	// A bidirectional Function
	ch := make(chan int)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func buffered() {
	wg := &sync.WaitGroup{}
	// One buffer for waiting
	ch := make(chan int, 1)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		// That will not show any outcome because there is no recever
		ch <- 27
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func uni_directional() {
	wg := &sync.WaitGroup{}
	// One buffer for waiting
	ch := make(chan int, 1)

	wg.Add(2)

	// Send only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	// receive only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 900
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

// Will output 42 and 0
// Because channel putput its zero value
// where there is no sender left
func close_channel() {
	wg := &sync.WaitGroup{}
	// One buffer for waiting
	ch := make(chan int, 1)

	wg.Add(2)
	// Send only
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		close(ch)
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)
	// receive only
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func channel_if() {
	wg := &sync.WaitGroup{}
	// One buffer for waiting
	ch := make(chan int, 1)

	wg.Add(2)
	// Send only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// If there is no return from channel
		// or 0 comes from a closed channel
		// Now the channel will only print if there is a message
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)
	// receive only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 3
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

func channel_for() {
	wg := &sync.WaitGroup{}
	// One buffer for waiting
	ch := make(chan int, 1)

	wg.Add(2)
	// Send only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	// receive only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// Inform the for loop above that there will not be no channel
		// Thats how you will stop the iteration
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
