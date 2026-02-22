package btree

type Item interface {
	Less (than Item) bool
}

type items []Item
type children []*Node

type Node struct {
	items items
	children children
}

type BTree struct {
	root *Node
	degree int
}

func (tree *BTree) maxItems() int {
	return 2*tree.degree - 1
}

func (tree *BTree) minItems() int {
	return tree.degree - 1
}

func NewTree(degree int) *BTree {
	if(degree < 2) {
		degree = 2
	}
	return &BTree{
		root: nil,
		degree: degree,
	}
}

func NewNode() *Node {
	return &Node{
		items: make(items, 0),
		children: make(children, 0),
	}
}

