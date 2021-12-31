package main

import (
	"fmt"
	"strconv"
)

// package level variable
var globalVariable float64 = 1

var (
	doctorName        = "Alexander ABC"
	doctorSeasonId    = 1
	doctorSalaryInBtc = 0.445
)

/*
Naming convention:
 - lowercase first letter for package scope
 - uppercase first letter to export, for example strconv.Itoa
 - no private scope
 - use PascalCase or camelCase, and capital	ize acronyms (HTTP, URL)
 - as short as reasonable - longer names for longer lives
*/

var (
	Count = 1
)

func main() {
	// three ways to declare variables
	var numberOfStudent int = 12
	var age = "data"
	floatNumber := float64(2) / 3
	fmt.Printf("%T %v \n%T %v\n%T %v\n", age, age, numberOfStudent, numberOfStudent, floatNumber, floatNumber)

	// shadowing, the innermost scope actually takes precedence
	fmt.Printf("Before shadowing doctor %T %v\n", doctorName, doctorName)
	var doctorName int = 1
	fmt.Printf("After shadowing doctorName %T %v\n", doctorName, doctorName)

	// In go, string is an alias for stream of bytes. So what happens when we asked the golang to convert to string
	// is it looks for what UNICODE CHARACTER is SET at the value 42, which is a star/asterisk
	var doctorSeasonId = 42
	var convertedDoctorSeasonId = string(doctorSeasonId)
	fmt.Printf("After shadowing doctorSeasonId:s %T %v\n", convertedDoctorSeasonId, convertedDoctorSeasonId)
	// right way to convert from numbers to the string
	convertedDoctorSeasonId = strconv.Itoa(doctorSeasonId)
	fmt.Printf("After shadowing doctorSeasonId at the right way: %T %v\n",
		convertedDoctorSeasonId, convertedDoctorSeasonId)

}
