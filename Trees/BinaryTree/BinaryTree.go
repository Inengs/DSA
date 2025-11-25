package BinaryTree

import "fmt"

type Node struct { // this defines the node struct for each node, containing the data, then the address of the left and right node
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node // this pointer stores the address of the root node whose data type is struct Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func CreateBT() *Node {
	var x int;
	// var newNode *Node
	fmt.Print("Enter data(-1 for no node): ") // ask for input from user
	fmt.Scan(&x)
	if (x == -1) {
		return nil // if input is -1, it means no node.
	}

	// Create Node with value
	newNode := &Node{Value: x}

	// Recursively create left subtree
	fmt.Printf("Enter left child of %d\n ", x)
	newNode.Left = CreateBT() // assign to left field,

	// Recursively create right subtree
	fmt.Printf("Enter right child of %d ", x)
	newNode = &Node{Right: CreateBT()} // assign to right field

	return newNode
}