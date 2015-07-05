package huffman

import ("tree"
		"os")

//Método para escrever a arvore recursivamente
func writeTree(node *tree.Node, writer *BitWriter) {
	if  node.isLeaf() { // folha
		writer.WriteBit(true)
		writer.WriteByte(([]byte(node.value))[0])
	} else { // tem dois filhos
		writer.WriteBit(false)
		writeNode(node.left, writer)
		writeNode(node.right, writer)
	}
}

// Método para criar um dicionario para o caracter e seu codigo gerado pelo algoritimo de Huffman

func createDict(node *tree.Node, dict *map[string]string, code string) {
	if node.isLeaf() {
		dict[node.value] = code
	} else {
		createDict(node.left, dict, code+"0")
		createDict(node.left, dict, code+"1")
	}
}

// Método para escrever o arquivo na forma codificada

func writeCodified(file os.File, dict *map[string]string, writer *BitWriter){
	for {
		b := file.Read(make([]byte, 1))
		//To-do: Escrever bit a bit da string que so tem 0s e 1s
		for {

		}
		writer.WriteBit(dict[string(b)])
	}

}

//Recebe um arquivo de texto e cria um arquivo comprimido
func Compress(file os.File, outputName string) {
	// gerar a arvore a partir da frequencia dos caracteres do texto
	root := harverst(getMap(os.File))

	// gerar dicionario
	dict := createDict(root, &make(map[string]string), "")

	//Resetar cursor
	file.Seek(0, 0)
	//Escrever Árvore
	writer = BitWriter.New()
	writeTree(root, writer)

	// Codificar



	return output
}

//Método para ler a arvore recursivamente
func readTree(reader *BitReader) *tree.Node{
	if reader.ReadBit() { // folha
		return tree.Node.New(string(reader.ReadByte()), nil, nil)
	} else { // tem dois filhos
		leftChild := readTree(reader)
		rightChild := readTree(reader)
		return tree.Node.New("", leftChild, rightChild)
	}
}

//Recebe um arquivo comprimido (objeto) e retorna o arquivo original (objeto)
func Decompress(file os.File, outputName string) os.File{
	// Ler Árvore (Reconstruir)
	reader = BitReader.New(os.File)
	root := readTree(reader)
	if root == nil {
		panic("Árvore nula!")
	}
	// Decodificar percorrendo a arvore
	output := os.Create(outputName)
	node := root
	for {
		bit, err := reader.ReadBit()
		if err != nil {
			break
		}
		// Anda na árvore, se bit=0 vai pro filho esquerdo
		if bit {
			node = node.right
		} else {
			node = node.left
		}

		//Checar se chegamos em uma folha
		if node.isLeaf() {
			output.WriteString(node.value)
			node := root
		}
	}

	return output
}

