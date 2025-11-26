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
	var x int
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
	newNode.Right=CreateBT() // assign to right field

	return newNode
}

func Preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Value)
	Preorder(root.Left)
	Preorder(root.Right)
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Printf("%d ", root.Value)
	Inorder(root.Right)
}

func Postorder(root *Node) {
	if root == nil {
		return
	}
	Postorder(root.Left)
	Postorder(root.Right)
	fmt.Printf("%d ", root.Value)
}

// InsertNode is a recursive helper method that inserts a value into the BST
// starting from the given current node
// Parameters:
//   - current: the node we're currently examining
//   - value: the integer value to insert
func (bt *BinaryTree) InsertNode(current *Node, value int) {
	// Compare the value with current node's value to determine direction
	if value < current.Value {
		// Value is less than current, so it belongs in the LEFT subtree
		
		if current.Left == nil {
			// Left child is empty - we found the insertion spot!
			// Create a new node and attach it as the left child
			current.Left = &Node{Value: value}
		} else {
			// Left child exists - keep searching deeper in left subtree
			// Recursively call InsertNode on the left child
			bt.InsertNode(current.Left, value)
		}
	} else {
		// Value is greater than or equal to current, so it goes RIGHT
		// Note: This implementation doesn't allow duplicates (they go right)
		
		if current.Right == nil {
			// Right child is empty 
			// Create a new node and attach it as the right child
			current.Right = &Node{Value: value}
		} else {
			// Right child exists - keep searching deeper in right subtree
			// Recursively call InsertNode on the right child
			bt.InsertNode(current.Right, value)
		}
	}
}

// Insert adds a new value to the Binary Search Tree
// This is the main public method that users call
// Parameters:
//   - value: the integer value to insert into the tree
func (bt *BinaryTree) Insert(value int) {
	// Check if the tree is empty (no root node exists)
	if bt.Root == nil {
		// Tree is empty - create the root node with the value
		bt.Root = &Node{Value: value}
	} else {
		// Tree already has nodes - use helper method to find insertion point
		// Start the recursive insertion from the root
		bt.InsertNode(bt.Root, value)
	}
}

// SEARCHING
func (bt *BinaryTree) Contains(root *Node, value int){

}