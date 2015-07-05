package tree

// Definição do nó da árvore de Huffman
type Node struct {
	value string
	left  *Node
	right *Node
}

func New(v string, l *Node, r *Node) *Node{
	node := new(Node)
	node.left = l
	node.right = r
	node.value = v
	return node
}

// Método que informa se um nó é uma folha
func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

