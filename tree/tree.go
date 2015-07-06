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

func (n *Node) String() (s string){
	var queueN []*Node{ n }
	var queueD []*int{ 0 }
	for len(queue) > 0 {
		// Tiramos o primeiro elemento e atualizamos a fila
		firstN := queueN[0]
		queueN := queueN[1:]

		firstD := queueD[0]
		queueD := queueD[1:]

		for i := 0;  i < firstD; i++ {
			s += "  "
		}
		s += "\"" + firstN.value + "\"" + "\n"
		if !firstN.isLeaf(){
			queueN.append(firstN.left)
			queueN.append(firstN.right)

			queueD.append(firstD + 1)
			queueD.append(firstD + 1)
		}
	}
	return
}
