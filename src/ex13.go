package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "***", i)
		time.Sleep(time.Second * 1)
	}
}

func say2(s string) {
	for i := 0; i <= 1000000; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// goroutine
	say("Sync")
	go say("Async1")
	go say("Async2")
	go say("Async3")
	time.Sleep(time.Second * 3)

	// anonymous goroutine
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()
	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")
	wait.Wait()

	// Concurrency vs Parallelism
	runtime.GOMAXPROCS(3)
	go say2("Async1")
	go say2("Async2")
	go say2("Async3")
	time.Sleep(time.Second * 12)

	// channel
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	println(i)

	// channel 2
	done := make(chan bool)
	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	<-done

	// channel deadlock
	// c := make(chan int)
	// c <- 1
	// fmt.Println(<-c)

	// buffer channel
	ch = make(chan int, 1)
	ch <- 101
	fmt.Println(<-ch)

	// channel sending/receiving
	chstr := make(chan string, 1)
	sendChnn(chstr)
	receiveChnn(chstr)

	// channel close
	ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	println(<-ch)
	println(<-ch)
	if _, success := <-ch; !success {
		println("no data")
	}

	// channel receiving 1
	ch = make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 3
	ch <- 2
	close(ch)
	for i := range ch {
		println(i)
	}

	// channel receiving 2
	ch = make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 1
	ch <- 4
	close(ch)
	for {
		if i, success := <-ch; success {
			println(i)
		} else {
			break
		}
	}

	// channel select
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1:
			println("run1 finished")
		case <-done2:
			println("run2 finished")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func sendChnn(ch chan<- string) {
	ch <- "Data"
}

func receiveChnn(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}
