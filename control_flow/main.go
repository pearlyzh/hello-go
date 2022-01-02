package main

import "fmt"

/*
 - defer: how can we actually invoke a function, but delay its execution to some future point in time
 - panic: the application comes into a state that it can no longer run anymore, and how go runtime can trigger that,
and how we can trigger that on our own.
 - recovery: recover from panic
*/

func main() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")
}
