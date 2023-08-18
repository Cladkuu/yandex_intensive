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
	depthObhod(n)
}

func depthObhod(n *node) {
	fmt.Println(n.value)

	for _, nodeValue := range n.next {
		depthObhod(nodeValue)
	}
}
