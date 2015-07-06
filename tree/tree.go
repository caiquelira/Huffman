package tree

// Definição do nó da árvore de Huffman
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func New(v string, l *Node, r *Node) *Node{
	node := new(Node)
	node.Left = l
	node.Right = r
	node.Value = v
	return node
}

// Método que informa se um nó é uma folha
func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
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
		s += "\"" + first.node.Value + "\"" + "\n"
		if !first.node.IsLeaf(){
			queue = append(queue, pair{depth: first.depth + 1, node: first.node.Left})
			queue = append(queue, pair{depth: first.depth + 1, node: first.node.Right})
		}
	}
	return
}
