package tree

// Definição do nó da árvore de Huffman
type Node struct {
	value string
	left  *Node
	right *Node
}

func (n *Node) New(v string, l *Node, r *Node) *Node{
	var res *Node
	res.left = l
	res.right = r
	res.value = v
	return res
}

// Método que informa se um nó é uma folha
func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

