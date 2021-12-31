package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Actor struct {
	name  string
	films []string
	node  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if i != size-1 {
				nextNode := queue[0]
				node.Next = nextNode
			}

			queue = append(queue, node.Left)
			queue = append(queue, node.Right)

		}
	}

	return root
}

func main() {
	myActor := Actor{
		name:  "james",
		films: []string{"shawshank redemption", "spider man"},
	}

	fmt.Println("myActor: ", myActor)

	anotherActor := &myActor
	anotherActor.name = "anotherJames"
	fmt.Println("myActor: ", myActor)

	// anonymous struct
	anonymousStruct := struct {
		name   string
		number int
	}{name: "nameee", number: 123}

	fmt.Println("anonymousStruct: ", anonymousStruct)
}
