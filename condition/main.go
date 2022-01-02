package main

import "fmt"

func main() {
	a := 10
	b := 20
	switch a {
	case 9:
		fmt.Println(a)
		break
	case 10:
		fmt.Println("Wrong")
		fallthrough
	default:
		fmt.Println("True")
		break
	}
	fmt.Println(b)
}
