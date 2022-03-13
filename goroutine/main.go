package main

import (
	"fmt"
	"sync"
	"time"
)

var i int
var mux sync.Mutex

// แล้วทำทุกอย่างผ่าน thing
type thing struct {
	i   int
	mux sync.Mutex
}

func main() {
	// wg := &sync.WaitGroup{}

	// // know number of go
	// wg.Add(3)
	// start := time.Now()
	// go doSomething(wg)
	// go doSomething(wg)
	// go doSomething(wg)

	// wg.Wait()
	// // time.Sleep(120 * time.Millisecond)
	// fmt.Println(time.Since(start))

	// // example()

	// reace condition
	// go func() {
	// 	for {
	// 		fmt.Println(read())
	// 	}
	// }()

	// for i := 0; i < 10; i++ {
	// 	write(i)
	// }

	// channel
	ch := make(chan int)
	qCh := make(chan struct{})

	go fibonacci(ch, qCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	qCh <- struct{}{}

	<-qCh
	fmt.Println("end")
}

// channel
func fibonacci(ch chan int, qCh chan struct{}) {
	a, b := 0, 1

	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <-qCh:
			qCh <- struct{}{}
			return
		}
	}
}

func write(n int) {
	mux.Lock()
	i = n
	mux.Unlock()
}

func read() int {
	mux.Lock()
	defer mux.Unlock()
	return i
}

func doSomething(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something")
	wg.Done()
}

func example() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Print("+")
		}
	}()
	time.Sleep(time.Second)
}
