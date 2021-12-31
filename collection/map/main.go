package main

import "fmt"

func main() {
	firstMap := map[string]string{"key1": "key1",
		"key2": "key3",
		"key3": "key3",
	}
	fmt.Printf("firstMap: %T - %v\n", firstMap["122"], firstMap)

	value, _ := firstMap["key"]
	_, found := firstMap["key"]
	fmt.Printf("firstMap: %T - %v\n", value == "", value == "")
	fmt.Printf("firstMap found?: %T - %v\n", found, found)

	// Map using make
	makeMap := make(map[int]string)
	makeMap[1] = "asd"
	makeMap[100] = "12123"
	delete(makeMap, 13)
	fmt.Printf("firstMap: %T - %v\n", makeMap, makeMap)
}
