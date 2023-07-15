package main

import (
	"fmt"

	"github.com/teooliver/go-algorithms/bTree"
)

func main() {
	tree := &bTree.Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)

	fmt.Println(tree.Search(400))
}
