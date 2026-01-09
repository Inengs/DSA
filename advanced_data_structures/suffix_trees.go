package advanceddatastructures

// BUILD_SUFFIX_TREE(S):
//     append '$' to S
//     root = new Node

//     for i from 0 to length(S)-1:
//         current = root
//         for j from i to length(S)-1:
//             if edge starting with S[j] exists:
//                 move along edge
//             else:
//                 create new edge/node

type Node struct {
	children map[rune]*Node
}

func NewNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}

type SuffixTree struct {
	root *Node
}

func NewSuffixTree() *SuffixTree {
	return &SuffixTree{root: NewNode()}
}

// Naive construction
func (st *SuffixTree) Build(s string) {
	s += "$"

	for i := 0; i < len(s); i++ {
		current := st.root
		for _, ch := range s[i:] {
			if _, exists := current.children[ch]; !exists {
				current.children[ch] = NewNode()
			}
			current = current.children[ch]
		}
	}
}

// Pattern search
func (st *SuffixTree) Search(pattern string) bool {
	current := st.root

	for _, ch := range pattern {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return true
}
