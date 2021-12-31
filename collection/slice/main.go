package main

import "fmt"

func main() {
	// array has a fixed size which has to know at compiled time so we better use slices
	// slices are naturally what are called referenced type
	// backed by array

	// BE AWARE THAT COPIES REFER TO THE SAME UNDERLYING ARRAY

	gradeSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	example1 := gradeSlice[1:4]
	fmt.Printf("example1: %v\n", example1)

	fmt.Printf("gradeSlice: %T - %v\n", gradeSlice, gradeSlice)

	newCutSlices := gradeSlice[len(gradeSlice)-1:]
	fmt.Printf("%T - %v\n", newCutSlices, newCutSlices)

	// remove element 3rd (3) from slices
	gradeSlice = append(gradeSlice[:3], gradeSlice[4:]...) // spread operation to spread things out
	fmt.Printf("After remove 3: %T - %v\n", gradeSlice, gradeSlice)

	fmt.Printf("example1 after gradeSlice removed 3: %v\n", example1)

	// use make to create a new slice
	sliceWithMake := make([]int, 0, 10)
	sliceWithMake = append(sliceWithMake, 1)
	fmt.Printf("sliceWithMake: %T - %v\n", sliceWithMake, sliceWithMake)
}
