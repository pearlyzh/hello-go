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

// & (address of): ampersand stands for is get the pointer
// *: the asterisk stands for dereference
// https://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/

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

	// Notes for array vs slices
	/*
			All assignments in Go are copy operations
			Slice is actually a projection of an underlying array, so the slice doesn't contain the data itself, the slice
			contains a pointer to the first element of the underlying array. So when we work with slices, the internal
			representation of a slice actually has an pointer to an array. So the data get copied is the pointer,
			not the data itself.
			The same with map. Slice and Map contain internal pointers, so copies point to the same underlying data

			If we are working with other data types, specifically PRIMITIVES, ARRAYS or STRUCTS. It's going to copy the entire
		 	structure unless you are using pointers.
	*/

	/*
			new(T) is equivalent to &T{} .
			make(T) : it returns an initialized value of type T ,
		It allocates and initializes the memory. Its used for slices, map and channels. You need make() to create channels and maps (and slices, but those can be created from arrays too).
	*/
}
