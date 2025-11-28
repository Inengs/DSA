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
// Contains checks if a value exists in the binary search tree
// root: the current node being examined (starts at tree root, changes with recursion)
// value: the integer value we're searching for
// returns: true if value is found, false otherwise
func (bt *BinaryTree) Contains(root *Node, value int) bool{
	// Base case: if current node is nil, we've reached the end without finding the value
	if root == nil {
		return false
	} 
	
	// Check if current node contains the value we're looking for
	if root.Value == value {
		return true // Found it!
	} else if value < root.Value {
		// Value is smaller than current node, so search in the left subtree
		return bt.Contains(root.Left, value)
	} else {
		// Value is larger than current node, so search in the right subtree
		return bt.Contains(root.Right, value)
	}
}

// Find a node
// Searches for a node with the given value in the BST
func (bt *BinaryTree) FindNode(root *Node, value int) *Node{
	// Base case: if we've gone past a leaf node, value doesn't exist
	if root == nil {
		return nil
	}
	// Found the node we're looking for
	if root.Value == value {
		return root
	} else if value < root.Value{
		// Value is smaller, search left subtree
		return bt.FindNode(root.Left, value)
	} else {
		// Value is larger, search right subtree
		return bt.FindNode(root.Right, value)
	}
}

// Find Parent
// Returns the parent node of the node containing the given value
func (bt *BinaryTree) FindParent(value int, root *Node) *Node{
	// If searching for root's value, it has no parent
	if value == root.Value {
		return nil
	} 
	
	// Search in left subtree
	if value < root.Value {
		// Left child doesn't exist
		if root.Left == nil {
			return nil
		// Left child is the node we're looking for - current node is its parent
		} else if root.Left.Value == value {
			return root
		// Keep searching deeper in left subtree
		} else {
			return bt.FindParent(value, root.Left)
		}
	// Search in right subtree
	} else {
		// Right child doesn't exist
		if root.Right == nil {
			return nil
		// Right child is the node we're looking for - current node is its parent
		} else if root.Right.Value == value {
			return root
		// Keep searching deeper in right subtree
		} else {
			return bt.FindParent(value, root.Right)
		} 
	}
}

// Size returns the total number of nodes in the tree
func (bt *BinaryTree) Size(node *Node) int {
	// Base case: empty tree/subtree has 0 nodes
    if node == nil { return 0 }
    // Count current node (1) + all nodes in left subtree + all nodes in right subtree
    return 1 + bt.Size(node.Left) + bt.Size(node.Right)
}

// DELETE
// Remove deletes a node with the given value from the BST
// Returns true if deletion was successful, false if value not found
func(bt *BinaryTree) Remove(value int) bool{
	// First, find the node we want to delete
	nodeToRemove := bt.FindNode(bt.Root, value)
	// Value doesn't exist in tree
	if nodeToRemove == nil {
		return false
	}
	
	// Find the parent of the node to remove (needed for reconnecting tree)
	parent := bt.FindParent(value, bt.Root)
	
	// CASE 1: Removing the only node in the tree (the root)
	if parent == nil && nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		bt.Root = nil
		return true
	}
	
	// CASE 2: Node is a leaf (no children)
	if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		// Disconnect this leaf from its parent
		if nodeToRemove.Value < parent.Value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	// CASE 3: Node has only a right child
	} else if nodeToRemove.Left == nil && nodeToRemove.Right != nil {
		// If removing root, update root pointer
		if parent == nil {
			bt.Root = nodeToRemove.Right
		// Otherwise, link parent to the node's right child
		} else if nodeToRemove.Value < parent.Value {
			parent.Left = nodeToRemove.Right
		} else {
			parent.Right = nodeToRemove.Right
		}
	// CASE 4: Node has only a left child
	} else if nodeToRemove.Left != nil && nodeToRemove.Right == nil {
		// If removing root, update root pointer
		if parent == nil {
			bt.Root = nodeToRemove.Left
		// Otherwise, link parent to the node's left child
		} else if nodeToRemove.Value < parent.Value {
			parent.Left = nodeToRemove.Left
		} else {
			parent.Right = nodeToRemove.Left
		}
	// CASE 5: Node has two children (most complex case)
	} else {
		// Find the in-order predecessor (largest value in left subtree)
		// This is the rightmost node in the left subtree
		largestValue := nodeToRemove.Left
		for largestValue.Right != nil {
			largestValue = largestValue.Right
		}
		
		// Find parent of the predecessor to disconnect it
		predecessorParent := bt.FindParent(largestValue.Value, bt.Root)
		
		// Handle case where predecessor is direct left child (has no right children)
		if predecessorParent == nodeToRemove {
			// The left child becomes the predecessor's left child
			nodeToRemove.Value = largestValue.Value
			nodeToRemove.Left = largestValue.Left
		} else {
			// Disconnect predecessor from its parent
			// Predecessor may have a left child, so connect it to predecessor's parent
			predecessorParent.Right = largestValue.Left
			// Replace the value of nodeToRemove with predecessor's value
			nodeToRemove.Value = largestValue.Value
		}
	}
	
	// Deletion successful
	return true
}
// FindMin returns the smallest value in the BST
// root: the current node being examined (start with tree root)
// In a BST, the minimum value is always the leftmost node
func (bt *BinaryTree) FindMin(root *Node) int{
	// Base case: if there's no left child, current node has the minimum value
	if root.Left == nil {
		return root.Value
	}
	// Recursive case: keep going left to find smaller values
	return bt.FindMin(root.Left)
}

// FindMax returns the largest value in the BST
// root: the current node being examined (start with tree root)
// In a BST, the maximum value is always the rightmost node
func (bt *BinaryTree) FindMax(root *Node) int{
	// Base case: if there's no right child, current node has the maximum value
	if root.Right == nil {
		return root.Value
	} 
	// Recursive case: keep going right to find larger values
	return bt.FindMax(root.Right)
}

// BreadthFirstTraversal performs level-order traversal (BFS)
func (bt *BinaryTree) BreadthFirstTraversal() {
	// Empty tree check
	if bt.Root == nil {
		return
	}
	
	// Create a queue to store nodes to visit
	queue := []*Node{bt.Root}
	
	// Process nodes level by level
	for len(queue) > 0 {
		// Dequeue the front node
		current := queue[0]
		queue = queue[1:]
		
		// Process current node
		fmt.Printf("%d ", current.Value)
		
		// Enqueue left child if it exists
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		
		// Enqueue right child if it exists
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}