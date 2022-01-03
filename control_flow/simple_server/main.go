package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hi there from Go!"))
		log.Println(ioutil.ReadAll(request.Body))
	})

	err := http.ListenAndServe("google.com:8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
