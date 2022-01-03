package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	value, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", value)
}
