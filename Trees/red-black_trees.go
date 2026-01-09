package trees

// insert(key):
//     node = BST_insert(key)
//     node.color = RED

//     while node != root and node.parent.color == RED:
//         if parent is left child of grandparent:
//             uncle = grandparent.right

//             if uncle.color == RED:
//                 parent.color = BLACK
//                 uncle.color = BLACK
//                 grandparent.color = RED
//                 node = grandparent
//             else:
//                 if node is right child:
//                     node = parent
//                     leftRotate(node)
//                 parent.color = BLACK
//                 grandparent.color = RED
//                 rightRotate(grandparent)
//         else:
//             symmetric cases

//     root.color = BLACK

import "fmt"

const (
	RED   = true
	BLACK = false
)

type Node struct {
	key    int
	color  bool
	left   *Node
	right  *Node
	parent *Node
}

type RedBlackTree struct {
	root *Node
}

func rbLeftRotate(t *RedBlackTree, x *Node) {
	y := x.right
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func rbRightRotate(t *RedBlackTree, y *Node) {
	x := y.left
	y.left = x.right

	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}

	x.right = y
	y.parent = x
}

func rbInsert(t *RedBlackTree, key int) {
	node := &Node{key: key, color: RED}

	var parent *Node
	current := t.root

	for current != nil {
		parent = current
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	node.parent = parent
	if parent == nil {
		t.root = node
	} else if key < parent.key {
		parent.left = node
	} else {
		parent.right = node
	}

	fixInsert(t, node)
}

func fixInsert(t *RedBlackTree, node *Node) {
	for node != t.root && node.parent.color == RED {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right

			if uncle != nil && uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					rbLeftRotate(t, node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				rbRightRotate(t, node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.left

			if uncle != nil && uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					rbRightRotate(t, node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				rbLeftRotate(t, node.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func rbInorder(n *Node) {
	if n != nil {
		rbInorder(n.left)
		fmt.Print(n.key, " ")
		rbInorder(n.right)
	}
}
