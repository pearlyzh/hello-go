package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string
	age  int
}

var wg = sync.WaitGroup{}

func main() {
	var channel = make(chan Person)
	wg.Add(2)
	go consume(channel)
	go produce(channel)
	wg.Wait()

	// ch := make(chan int)
	// wg.Add(2)
	// go func() {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	// for j := 0; j < 5; j++ {
	// 	wg.Add(2)
	// 	go func() {
	// 		ch <- 42
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
}

func consume(channel chan Person) {
	start := time.Now()
	data := <-channel
	data = <-channel
	elapsed := time.Since(start)
	fmt.Printf("Hey, I have consumed this %v after waiting time %v\n", data, elapsed)
	wg.Done()
}

func produce(channel chan Person) {
	time.Sleep(3 * time.Second)
	data := new(Person)
	data.age = 10
	data.name = "pearly"
	fmt.Println("I am going to send this data: ", *data)
	channel <- *data
	data.name = "something else"
	data.age = 100
	fmt.Println("I have changed the data: ", *data)
	wg.Done()
}
