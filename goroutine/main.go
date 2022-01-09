package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Inside of Go runtime, we've got a scheduler which is going to map these goroutines onto these operating system threads
for period of time, and the scheduler with then take turns with every CPU thread that's available and assign the different
goroutines, a certain amount of processing time on those threads.
*/

/*
- Don't create goroutine in libraries
	- Let the consumers control the concurrency
- When creating goroutines,know how they will end
 	- Avoid subtle memory leaks
- Check race conditions at compiled time
*/

/*
 - Synchronization:
	- Use WaitGroup to wait for groups of goroutines to complete
	- Use Mutex and RWMutex to protect data access
 - Parallelism:
	- By default, Go will use CPU threads equal to the number of available CPU cores
	- Change with GOMAXPROCS
	- More thread can improve performance,but too many can slow it down
*/

func main() {
	// waitGroup := sync.WaitGroup{}
	// waitGroup.Add(1)
	// go sayHello(&waitGroup)
	// time.Sleep(5 * time.Second)
	// fmt.Println("This is a line from main!")
	// waitGroup.Wait()

	fmt.Println(runtime.GOMAXPROCS(-1))

	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		mutex.RLock()
		go sayHello()
		go increment()
	}
	waitGroup.Wait()
}

var counter = 0
var waitGroup = sync.WaitGroup{}
var mutex = sync.RWMutex{}

func increment() {
	// mutex.Lock()
	counter++
	// mutex.Unlock()
	waitGroup.Done()
}

func sayHello() {
	// mutex.RLock()
	fmt.Println("I am saying hello and the counter is ", counter)
	mutex.RUnlock()
	waitGroup.Done()
}
