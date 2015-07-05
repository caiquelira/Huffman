package huffman

import ("github.com/caiquelira/huffman/tree"
		"github.com/caiquelira/huffman/bitreader"
		"github.com/caiquelira/huffman/bitwriter"
		"os")

//Método para escrever a arvore recursivamente
func writeTree(node *tree.Node, writer *bitwriter.BitWriter) {
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

func createDict(node *tree.Node, dict map[string]string, code string) {
	if node.isLeaf() {
		dict[node.value] = code
	} else {
		createDict(node.left, dict, code+"0")
		createDict(node.left, dict, code+"1")
	}
}

// Método para escrever o arquivo na forma codificada

func writeCodified(file os.File, dict map[string]string, writer *bitwriter.BitWriter){
	//Loop para ler um caracter e escreve-lo no arquivo de saida em forma codificada
	for {
		b := file.Read(make([]byte, 1))
		//Transformar o caracter lido no código feito pelo dicionário
		codeb := dict[string(b)]
		//Temos que escrever bit a bit.
		for i := 0; i < len(codeb); i++ {
			if codeb[i] == "1"{
				writer.WriteBit(true)
			} else {
				writer.WriteBit(false)
			}

		}
	}
	writer.Close()
}

//Recebe um arquivo de texto e cria um arquivo comprimido
func Compress(file os.File, outputName string) {
	// gerar a arvore a partir da frequencia dos caracteres do texto
	root := harverst(getMap(os.File))

	// gerar dicionario
	dict := make(map[string]string)
	createDict(root, dict, "")

	//Resetar cursor
	file.Seek(0, 0)
	//Escrever Árvore
	writer = bitwriter.New()
	writeTree(root, writer)

	// Codificar

	writeCodified(file, dict, writer)
}

//Método para ler a arvore recursivamente
func readTree(reader *bitreader.BitReader) *tree.Node{
	if reader.ReadBit() { // folha
		return tree.Node.New(string(reader.ReadByte()), nil, nil)
	} else { // tem dois filhos
		leftChild := readTree(reader)
		rightChild := readTree(reader)
		return tree.Node.New("", leftChild, rightChild)
	}
}

func decodeFile(reader *bitreader.BitReader, outputName string, root *tree.Node) os.File {
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
//Recebe um arquivo comprimido (objeto) e retorna o arquivo original (objeto)
func Decompress(file os.File, outputName string) os.File{
	// Ler Árvore (Reconstruir)
	reader = bitreader.New(os.File)
	root := readTree(reader)
	if root == nil {
		panic("Árvore nula!")
	}
	// Decodificar percorrendo a arvore

	return decodeFile(reader, outputName, root)
}

