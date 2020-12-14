package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode

}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("node id null")
		return
	}
	node.value = value
}

func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

func (node *treeNode) rangeNode() {     
	if node == nil {
		return
	}
	node.left.rangeNode()
	node.print()
	node.right.rangeNode()
}

func main() {
	var root treeNode
	root.value = 10
	fmt.Println(root)
	root.left.setValue(9)
	//root.right.setValue(8)
	//root.left.left.setValue(7)
	//root.left.right.setValue(6)
	//root.right.left.setValue(5)
	//root.right.right.setValue(4)
	root.rangeNode()
}
