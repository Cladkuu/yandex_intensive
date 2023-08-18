package main

import "fmt"

type node struct {
	value int
	next  []*node
}

func main() {
	n := &node{value: 1, next: []*node{
		{value: 2,
			next: []*node{
				{
					value: 3,
				},
				{
					value: 115,
				},
			}},
		{
			value: 4,
			next: []*node{
				{
					value: 232,
				},
			},
		},
		{
			value: 6,
			next: []*node{
				{
					value: 7,
				},
			},
		},
	}}

	widthObhod(n)
}

func widthObhod(n *node) {
	fmt.Println(n.value)
	queue := make([]*node, 0, 10)
	queue = append(queue, n.next...)

	var temp []*node
	for len(queue) > 0 {
		temp = queue
		for _, val := range temp {
			fmt.Println(val.value)
			queue = queue[1:]
			queue = append(queue, val.next...)
		}
	}
}
