package tree

type Node struct {
	Value       int
	Left, Right *Node
}

func NewNode(value int) *Node {
	node := &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
	return node
}

func (n *Node) InsertLeft(value int) *Node {
	left := NewNode(value)
	n.Left = left
	return left
}

func (n *Node) InsertRight(value int) *Node {
	right := NewNode(value)
	n.Right = right
	return right
}

func CreateTree(i int, values map[int]int) *Node {
	val, isPresent := values[i]
	if i <= len(values) && isPresent {
		root := &Node{
			Value: val,
			Left:  CreateTree(2*i+1, values),
			Right: CreateTree(2*i+2, values),
		}
		return root
	}
	return nil
}
