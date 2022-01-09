package main

import (
	"fmt"
)

/*
Best practices:
- Use many, small interfaces
	- Single method interfaces are some of the most powerful and flexible
		-io.Writer, io.Reader, interface{}
- Don't export interfaces for types that will be consumed
- Do export interfaces for types that will be used by package
*/

func main() {
	var writer Writer = ConsoleWriter{}
	writer.Write([]byte("nasdsa"))
}

// Writer interface does not describe data, interface describes behaviour
type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
