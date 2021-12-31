package main

import "fmt"

/* Summary
- Immutable, but can be shadowed
- Replaced by the compiler at compile time

*/

// enumerated constants
const (
	// a counter we can use
	a = iota
	// b is applied the same pattern and use iota as well
	b
)

// iota is reset to 0 because now we have another constant block
const (
	anotherConstant = iota
)

const (
	// underscore: tells the compiler: I know you gonna generate a value here,
	// but I don't care what it is and just throw it away
	_ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

// use as bit shifting
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	isAdmin = 1 << iota
	isHeadquarter
	isOperator

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeArctic
	canSeeAntarctic
)

func main() {
	// we don't want to export it
	const myConst = "bitch"

	// we want to export it
	const MySecondConst = "bitch"

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", anotherConstant, anotherConstant)

	var mySpecialist = catSpecialist
	myStrangeSpecialist := 0
	fmt.Printf("%v\n", mySpecialist == catSpecialist)
	fmt.Printf("%v - %v\n", myStrangeSpecialist == catSpecialist, catSpecialist)

	fileSizeInBytes := 1048576. * 1024
	fmt.Printf("%T - %v (GB)\n", fileSizeInBytes/GB, fileSizeInBytes/GB)

	roles := isAdmin | canSeeAfrica | canSeeAsia
	fmt.Printf("%T - %b\n", roles, roles)
	fmt.Printf("isAdmin? - %v\n", roles&isAdmin == isAdmin)
	fmt.Printf("canSeeAsia? - %v\n", roles&canSeeAsia == canSeeAsia)
	fmt.Printf("canSeeArctic? - %v\n", roles&canSeeArctic == canSeeArctic)

	// conversion
	const myStrangeConst = 10
	var try1 int32 = 20
	var try2 byte = 5
	fmt.Printf("%T - %v\n", try1+myStrangeConst, try1+myStrangeConst)
	fmt.Printf("%T - %v\n", try2+myStrangeConst, try2+myStrangeConst)
}
