package trees

// insert(node, key):
//     if node == null:
//         return new Node(key)

//     if key < node.key:
//         node.left = insert(node.left, key)
//     else if key > node.key:
//         node.right = insert(node.right, key)
//     else:
//         return node

//     update height of node
//     balance = height(left) - height(right)

//     if balance > 1 and key < node.left.key:
//         return rightRotate(node)

//     if balance < -1 and key > node.right.key:
//         return leftRotate(node)

//     if balance > 1 and key > node.left.key:
//         node.left = leftRotate(node.left)
//         return rightRotate(node)

//     if balance < -1 and key < node.right.key:
//         node.right = rightRotate(node.right)
//         return leftRotate(node)

//     return node

import "fmt"

type avlNode struct {
	key    int
	height int
	left   *avlNode
	right  *avlNode
}

func height(n *avlNode) int {
	if n == nil {
		return 0
	}
	return n.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rightRotate(y *avlNode) *avlNode {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func leftRotate(x *avlNode) *avlNode {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

func getBalance(n *avlNode) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func insert(node *avlNode, key int) *avlNode {
	if node == nil {
		return &avlNode{key: key, height: 1}
	}

	if key < node.key {
		node.left = insert(node.left, key)
	} else if key > node.key {
		node.right = insert(node.right, key)
	} else {
		return node
	}

	node.height = max(height(node.left), height(node.right)) + 1
	balance := getBalance(node)

	// LL
	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	// RR
	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	// LR
	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// RL
	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func inorder(root *avlNode) {
	if root != nil {
		inorder(root.left)
		fmt.Print(root.key, " ")
		inorder(root.right)
	}
}