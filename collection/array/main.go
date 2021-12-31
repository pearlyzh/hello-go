package main

import "fmt"

func main() {
	// array has a fixed size which has to know at compiled time so we better use slices
	grades := [...]int{1, 2, 3, 4}
	var students []string
	students = append(students, "132")

	fmt.Printf("Grades %T-%v\n", grades, grades)

	fmt.Printf("Students %T-%v\n", students, students)

	myCopiedGrades := &grades
	myCopiedGrades[2] = 1000
	myCopiedGrades[3] = 1001
	fmt.Printf("Grades Again %T-%v\n", grades, grades)
	fmt.Printf("MyCopiedGrades %T-%v\n", myCopiedGrades, myCopiedGrades)

	// 2 dimension array
	var myTwoIntDimension = [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	for _, nums := range myTwoIntDimension {
		for _, num := range nums {
			fmt.Printf("%v - ", num)
		}
		fmt.Printf("\n")
	}
}
