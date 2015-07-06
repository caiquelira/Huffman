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

type pair struct{
	depth int
	node *Node
}

func (n *Node) String() (s string){
	queue := []pair{ pair{depth: 0, node: n,} }

	for len(queue) > 0 {
		// Tiramos o primeiro elemento e atualizamos a fila
		first := queue[0]
		queue := queue[1:]

		for i := 0;  i < first.depth; i++ {
			s += "  "
		}
		s += "\"" + first.node.value + "\"" + "\n"
		if !first.node.isLeaf(){
			queue = append(queue, pair{depth: first.depth + 1, node: first.node.left})
			queue = append(queue, pair{depth: first.depth + 1, node: first.node.right})
		}
	}
	return
}
