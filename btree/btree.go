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

//Btree methods
func (tree *BTree) maxItems() int {
	return 2*tree.degree - 1
}

func (tree *BTree) minItems() int {
	return tree.degree - 1
}

func (tree *BTree) Insert(item Item){
	if(tree.root == nil) {
		tree.root = NewNode()
		tree.root.items = append(tree.root.items, item)
	} else if len(tree.root.children) == 0 {
		if len(tree.root.items) < tree.maxItems(){
			tree.root.items.insert(item)
		} else {
			newRoot := NewNode()
			newRoot.children = append(newRoot.children, tree.root)
			left, right := tree.root.splitNode(tree.minItems()-1)
			newRoot.children[0] = left
			newRoot.children = append(newRoot.children, right)
			newRoot.items.insertAt(0, tree.root.items[tree.minItems()])
			tree.root = newRoot
			
		}
	} else {
		tree.root.insert(item)
	}
}



//Node methods
func (node *Node) find(item Item) (int, bool) {
	n := len(node.items)
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if node.items[mid].Less(item) {
			low = mid + 1
		} else if item.Less(node.items[mid]) {
			high = mid - 1
		} else {
			return mid, true
		}
	}
	return low, false
}

func (node *Node) splitNode(splitAt int) (*Node, *Node){
	right := NewNode()
	right.items = append(right.items, node.items[splitAt+1:]...)
	node.items = node.items[:splitAt]
	return node, right
}

func (node *Node) insert(item Item) {
	index, found := node.find(item)
	if found {
		node.items[index] = item
		return
	}

	if len(node.children) == 0 {
		node.items.insertAt(index, item)
		return
	}

	if node.items[index].Less(item) {
	
	} else {
		
	}
}


//Children methods
func (children *children) insert(index int, node *Node){
	*children = append(*children, nil)
	copy((*children)[index+1:], (*children)[index:])
	(*children)[index] = node
}




//Items methods
func (items *items) insert(item Item) {

}

func (items *items) insertAt(index int, item Item) {
	*items = append(*items, nil)
	copy((*items)[index+1:], (*items)[index:])
	(*items)[index] = item
}