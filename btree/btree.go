package btree

type Item interface {
	Less(than Item) bool
}

type items []Item
type children []*Node

type Node struct {
	items    items
	children children
}

type BTree struct {
	root   *Node
	degree int
}

func NewTree(degree int) *BTree {
	if degree < 2 {
		degree = 2
	}
	return &BTree{
		root:   nil,
		degree: degree,
	}
}

func NewNode() *Node {
	return &Node{
		items:    make(items, 0),
		children: make(children, 0),
	}
}

// Btree methods
func (tree *BTree) maxItems() int {
	return 2*tree.degree - 1
}

func (tree *BTree) minItems() int {
	return tree.degree - 1
}

func (tree *BTree) Insert(item Item) {
	if tree.root == nil {
		tree.root = NewNode()
		tree.root.items = append(tree.root.items, item)
		return
	} else if len(tree.root.items) == tree.maxItems() {
		newRoot := NewNode()
		newRoot.children = append(newRoot.children, tree.root)
		promoted, right := tree.root.splitNode(tree.minItems())
		newRoot.items.insertAt(0, promoted)
		newRoot.children = append(newRoot.children, right)
		tree.root = newRoot
	}
	tree.root.insert(item, tree.degree)
}

//Node methods

func (node *Node) splitNode(splitAt int) (Item, *Node) {
	promoted := node.items[splitAt]
	right := NewNode()
	right.items = append(right.items, node.items[splitAt+1:]...)
	node.items = node.items[:splitAt]
	if(len(node.children) > 0) {
		right.children = append(right.children, node.children[splitAt+1:]...)
		node.children = node.children[:splitAt+1]
	}
	return promoted , right
}

func (node *Node) insert(item Item, degree int) {
	index, found := node.items.find(item)
	if found {
		node.items[index] = item
		return
	}

	if len(node.children) == 0 {
		node.items.insertAt(index, item)
		return
	}

	if len(node.children[index].items) == 2*degree-1 {
		node.items.insertAt(index, node.children[index].items[degree-1])
		_, right := node.children[index].splitNode(degree - 1)
		node.children.insertAt(index+1, right)
		if item.Less(node.items[index]) {
			node.children[index].insert(item, degree)
		} else {
			node.children[index+1].insert(item, degree)
		}
	} else {
		node.children[index].insert(item, degree)
	}
}

// Children methods
func (children *children) insertAt(index int, node *Node) {
	*children = append(*children, nil)
	copy((*children)[index+1:], (*children)[index:])
	(*children)[index] = node
}


// Items methods
func (items items) find(item Item) (int, bool) {
	n := len(items)
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if items[mid].Less(item) {
			low = mid + 1
		} else if item.Less(items[mid]) {
			high = mid - 1
		} else {
			return mid, true
		}
	}
	return low, false
}

func (items *items) insert(item Item) {
	index, found := items.find(item)
	if found {
		(*items)[index] = item
		return
	}
	items.insertAt(index, item)
}

func (items *items) insertAt(index int, item Item) {
	*items = append(*items, nil)
	copy((*items)[index+1:], (*items)[index:])
	(*items)[index] = item
}
