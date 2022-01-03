package curl

import "fmt"

/*
 - defer: how can we actually invoke a function, but delay its execution to some future point in time
 - panic: the application comes into a state that it can no longer run anymore, and how go runtime can trigger that,
and how we can trigger that on our own.
 - recovery: recover from panic
*/

func main() {
	fmt.Println("start")
	// LIFO order of defer so that we can close the resources in the opposite way when we open them
	defer fmt.Println("defer2")
	// same like finally in java, move this line AFTER the main function, and BEFORE
	// the main function return
	defer fmt.Println("defer1")
	fmt.Println("end")

	func() {

	}()
}
