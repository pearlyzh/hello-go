package main

import (
	"fmt"
	"strconv"
)

func main() {
	// default value is 0, which is false in bool
	// not an alias for other types
	var myPrimitive bool
	fmt.Printf("%v, %T\n", myPrimitive, myPrimitive)

	// int type, default type would be int
	var myInt = 42
	fmt.Printf("%v, %T\n", myInt, myInt)

	// int type, default type would be int
	var myUint uint8 = uint8(myInt * -200)
	fmt.Printf("%v, %T\n", myUint, myUint)

	// cannot use arithmetic operators on different types
	fmt.Printf("%v\n", uint16(myUint)+uint16(myInt))

	// bit operator
	a := 10                  // 1010
	b := 3                   // 0011 - NOT (1100)
	fmt.Printf("%v\n", a&b)  // 0010
	fmt.Printf("%v\n", a|b)  // 1011
	fmt.Printf("%v\n", a^b)  // 1001
	fmt.Printf("%v\n", a&^b) // a AND NOT(b) = 1000
	fmt.Printf("%v\n", a<<3) // 10 * 2^3 = 80
	fmt.Printf("%v\n", a>>2) // 10 / 2^2 = 2

	// float types
	// follow IEEE-754 standard
	// 32 bit version and64 bit version
	// flashback:https://towardsdatascience.com/binary-representation-of-the-floating-point-numbers-77d7364723f1+&cd=11
	myFloat32 := 0.1 * 3
	fmt.Printf("%v, %T\n", myFloat32, myFloat32)
	fmt.Printf("%v, %T\n", myFloat32 == 0.3, myFloat32 == 0.3)

	// complex type
	myComplex := 1.1e3 + 2.44e2i
	myAnotherComplex := complex(3.1e3, 2.2e1)
	fmt.Printf("%v, %T\n", myComplex*myAnotherComplex, myComplex)
	fmt.Printf("%v, %T\n", real(myComplex), real(myComplex))
	fmt.Printf("%v, %T\n", imag(myComplex), imag(myComplex))

	// string
	// string are generally immutable. Ref: https://www.baeldung.com/java-string-immutable
	// UTF-8, can be concatenated with + operator
	// can be converted into []byte
	myString := "this is a my string"
	// a lot of functions that we are going to use in go will work with bytes slices, like invoking service calls, sending files...
	myBytesSlices := []byte(myString)
	fmt.Printf("%v, %T\n", myString, myString)
	fmt.Printf("%v, %T\n", string(myString[2]), myString[2])
	fmt.Printf("%v, %T\n", myBytesSlices, myBytesSlices)

	// rune
	// alias for int32
	var myRune = 'A'
	fmt.Printf("%v, %T\n", myRune, myRune)
	fmt.Printf("%v, %T\n", string(myRune), string(myRune))
	fmt.Printf("%v, %T\n", strconv.QuoteRune(myRune), strconv.QuoteRune(myRune))

	myFunc := myFunc
	fmt.Printf("%T\n", myFunc)
}

func myFunc() {

}
