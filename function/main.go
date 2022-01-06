package main

import "fmt"

// https://stackoverflow.com/questions/8263546/whats-the-difference-of-functions-and-methods-in-go
// functions in go can be treated as types as well, they can be passed around like variables
// slice and map are always passed by value of pointers

type Person struct {
	firstName string
	lastName  string
}

func changeName(p *Person) {
	p = &Person{firstName: "nghia", lastName: "truong"}
}

func main() {
	person := Person{
		firstName: "Alice",
		lastName:  "Dow",
	}
	fmt.Printf("beforeChangeName %p %v\n", &person, person)
	changeName(&person)
	fmt.Printf("afterChangeName %p %v\n", &person, person)

	myValue := "value"
	fmt.Println("beforeMyFunction", &myValue, myValue)
	myFunction(&myValue)
	fmt.Println("afterMyFunction", &myValue, myValue)

	result := partition("aab")
	fmt.Println("result", result)

	// anonymous function
	// functions don't have names if they are:
	// - immediately invoked
	// - assigned to a variable or passed as an argument to a function
	var myFunc func() = func() {
		fmt.Println("myFunc")
	}
	myFunc()

	// methods:
	// function that executes in context of a type
	// func(g greeter) greet() // g is called a receiver
}

type counter string

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot provide by zero")
	}
	return float64(a / b), nil
} /*
* (also known as indirection operator) is used to “dereference” pointer variables,
dereferencing a pointer gives us access to the value the pointer points to.
In this case: *myString = "anotherString"
basically you are saying to the compiler:
store the string "Docker Extended" in the memory location course refers to.
*/
func myFunction(myString *string) {
	*myString = "anotherString"
}

/* Variadic parameter: Go runtime takes the last arguments which are passed in, and wrap them up into a slice which
has a name of the variable that we have here
*/
func mSum(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}

	return sum
}

/*
This is just a random leetcode problem solved in go
*/
func partition(s string) [][]string {
	result := new([][]string)
	helper(result, s, *new([]string), 0)
	// can return address of local variable
	// automatically promoted from local memory (stack) to shared memory (heap)
	return *result
}

func helper(result *[][]string, s string, cur []string, index int) {
	if index == len(s) {
		cpy := make([]string, len(cur))
		copy(cpy, cur)
		*result = append(*result, cpy)
		return
	}

	for i := index; i < len(s); i++ {
		if isPalindrome(s, index, i) {
			cur = append(cur, s[index:i+1])
			helper(result, s, cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}
}

func isPalindrome(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func carPooling(trips [][]int, capacity int) bool {
	stops := make([]int, 1001)

	for _, trip := range trips {
		stops[trip[1]] += trip[0]
		stops[trip[2]] -= trip[0]
	}

	cap := 0
	for i := 0; i < 1001; i++ {
		cap += stops[i]
		if cap > capacity {
			return false
		}
	}

	return true
}
