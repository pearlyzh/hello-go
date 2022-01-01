package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	foo   int
	actor Actor
}

type Actor struct {
	name  string
	films []string
}

// &: ampersand stands for is get the pointer
// *: the asterisk stands for dereference
func main() {
	a := 10
	// this asterisk before TYPE (int) is declaring the pointer to data of that type,
	// so this is a pointer to an integer
	var b *int = &a
	fmt.Println(a, b)
	fmt.Printf("Value and type: %v - %T\n", b, b)

	// dereference an integer: put an asterisk in front of a pointer
	fmt.Printf("Value and type: %v - %T\n", *b, *b)

	myArray := [3]int{1, 2, 3}
	myArrayRef := &myArray // "myArrayRef" is holding the memory location which is holding "myArray"
	myArrayRef[2] = 20
	fmt.Println(myArray, *myArrayRef)
	fmt.Printf("Type of values in myArray: %v\n", unsafe.Sizeof(myArray[2]))

	intPointer0 := &myArray[0]
	intPointer1 := &myArray[1]

	fmt.Printf("myArray adress %p\n", intPointer0)
	fmt.Printf("myArray adress %p\n", intPointer1)
	fmt.Printf("Value and type and pointer adress: %v - %T - %p\n", *myArrayRef, *myArrayRef, myArrayRef)
	fmt.Printf("Value and type and pointer adress: %v - %T - %p\n", *myArrayRef, *myArrayRef, myArrayRef)

	// with struct
	myStruct := MyStruct{
		foo: 400,
	}

	myStructPointer := &myStruct
	// Unlike new, make's return type is the same as the type of its
	// argument, not a pointer to it
	// make can only be used to allocate and initialize data of slice, map, channel type
	sliceWithMake := make([]int, 10)
	fmt.Println("sliceWithMake: ", sliceWithMake)

	// The new built-in function allocates memory. The first argument is a type,
	// not a value, and the value returned is a pointer to a newly
	sliceWithNew := new([]int)
	fmt.Println("sliceWithNew: ", sliceWithNew)

	// output: myStructPointer:  &{400}
	// which is basically saying that "myStructPointer"
	// is holding the value of the address of the object containing an integer field with the value 400 in it
	// new can allocate any type of data.
	fmt.Println("myStructPointer: ", myStructPointer)

	myActorWithNew := new(Actor)
	fmt.Println("myActorWithNew: ", myActorWithNew)

	// nil pointer
	var myStructNil *MyStruct

	fmt.Println("myStructNil", myStructNil)
	myStructNil = new(MyStruct)
	myStructNil.foo = 100
	// would be the same
	// (*myStructNil).foo = 100
	fmt.Println("myStructNil.actor.name: ", (*myStructNil).foo)

	// compare a pointer to an int in an array to the pointer to the address of that array
	myAnotherArray := [3]int{1, 2, 3}
	myIntPointer := &myAnotherArray[0]
	myAnotherArrayPointer := &myAnotherArray

	fmt.Printf("myIntPointer: %v - %T - %p\n", myIntPointer, myIntPointer, myIntPointer)
	fmt.Printf("myAnotherArrayPointer: %v - %T - %p\n", myAnotherArrayPointer, myAnotherArrayPointer, myAnotherArrayPointer)
	fmt.Printf("%v\n", *myIntPointer == (*myAnotherArrayPointer)[0])
	// would be the same
	fmt.Printf("%v\n", *myIntPointer == myAnotherArrayPointer[0])
}
