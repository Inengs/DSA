package trees

// splay(root, x):
//     if root == null or root.key == x:
//         return root

//     if x < root.key:
//         if root.left == null:
//             return root

//         if x < root.left.key:
//             root.left.left = splay(root.left.left, x)
//             root = rightRotate(root)
//         else if x > root.left.key:
//             root.left.right = splay(root.left.right, x)
//             if root.left.right != null:
//                 root.left = leftRotate(root.left)

//         return (root.left == null) ? root : rightRotate(root)

//     else:
//         symmetric logic for right subtree

import "fmt"

type SplayNode struct {
	key   int
	left  *SplayNode
	right *SplayNode
}

func SplayrightRotate(x *SplayNode) *SplayNode {
	y := x.left
	x.left = y.right
	y.right = x
	return y
}

func SplayleftRotate(x *SplayNode) *SplayNode {
	y := x.right
	x.right = y.left
	y.left = x
	return y
}

func splay(root *SplayNode, key int) *SplayNode {
	if root == nil || root.key == key {
		return root
	}

	if key < root.key {
		if root.left == nil {
			return root
		}

		if key < root.left.key {
			root.left.left = splay(root.left.left, key)
			root = SplayrightRotate(root)
		} else if key > root.left.key {
			root.left.right = splay(root.left.right, key)
			if root.left.right != nil {
				root.left = SplayleftRotate(root.left)
			}
		}

		if root.left == nil {
			return root
		}
		return SplayrightRotate(root)
	} else {
		if root.right == nil {
			return root
		}

		if key > root.right.key {
			root.right.right = splay(root.right.right, key)
			root = SplayleftRotate(root)
		} else if key < root.right.key {
			root.right.left = splay(root.right.left, key)
			if root.right.left != nil {
				root.right = SplayrightRotate(root.right)
			}
		}

		if root.right == nil {
			return root
		}
		return SplayleftRotate(root)
	}
}

func Splayinsert(root *SplayNode, key int) *SplayNode {
	if root == nil {
		return &SplayNode{key: key}
	}

	root = splay(root, key)

	if root.key == key {
		return root
	}

	newNode := &SplayNode{key: key}
	if key < root.key {
		newNode.right = root
		newNode.left = root.left
		root.left = nil
	} else {
		newNode.left = root
		newNode.right = root.right
		root.right = nil
	}

	return newNode
}

func Splayinorder(root *SplayNode) {
	if root != nil {
		Splayinorder(root.left)
		fmt.Print(root.key, " ")
		Splayinorder(root.right)
	}
}
