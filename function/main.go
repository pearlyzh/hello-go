package main

import "fmt"

// https://stackoverflow.com/questions/8263546/whats-the-difference-of-functions-and-methods-in-go

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
}

/*
* (also know as indirection operator) is used to “dereference” pointer variables,
dereferencing a pointer gives us access to the value the pointer points to.
In this case: *myString = "anotherString"
basically you are saying to the compiler:
store the string "Docker Extended" in the memory location course refers to.
*/
func myFunction(myString *string) {
	*myString = "anotherString"
}

func partition(s string) [][]string {
	result := new([][]string)
	helper(result, s, *new([]string), 0)
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
