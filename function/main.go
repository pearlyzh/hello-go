package main

func myFunction(myString *string) string {
	myFuncString := *myString
	println(&myFuncString)
	return myFuncString
}

func main() {
	myValue := "value"
	myFunction(&myValue)
	println(&myValue)

}

func findJudge(n int, trusts [][]int) int {
	trusted := make([]int, n+1, n+1)
	for _, trust := range trusts {
		trusted[trust[0]]--
		trusted[trust[1]]++
	}

	for i := 1; i <= n; i++ {
		if trusted[i] == n-1 {
			return i
		}
	}

	return -1
}
