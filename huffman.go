package huffman

import ("github.com/caiquelira/huffman/tree"
		"github.com/caiquelira/huffman/bit"
		"os"
		"io")

//Método para escrever a arvore recursivamente
func writeNode(node *tree.Node, writer *bit.Writer) {
	if  node.IsLeaf() { // folha
		writer.Write(true)
		writer.WriteByte(([]byte(node.Value))[0])
	} else { // tem dois filhos
		writer.Write(false)
		writeNode(node.Left, writer)
		writeNode(node.Right, writer)
	}
}

// Método para criar um dicionario para o caracter e seu codigo gerado pelo algoritimo de Huffman

func createDict(node *tree.Node, dict map[string]string, code string) {
	if node.IsLeaf() {
		dict[node.Value] = code
	} else {
		createDict(node.Left, dict, code+"0")
		createDict(node.Right, dict, code+"1")
	}
}

// Método para escrever o arquivo na forma codificada

func writeCodified(file *os.File, dict map[string]string, writer *bit.Writer){
	//Loop para ler um caracter e escreve-lo no arquivo de saida em forma codificada
	for {
		b := make([]byte, 1)
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}
		//Transformar o caracter lido no código feito pelo dicionário
		codeb := dict[string(b)]
		//Temos que escrever bit a bit.
		for i := 0; i < len(codeb); i++ {
			if string(codeb[i]) == "1"{
				writer.Write(true)
			} else {
				writer.Write(false)
			}
		}
	}
	writer.Close()
}

//Recebe um arquivo de texto e cria um arquivo comprimido
func Compress(file *os.File, outputName string) {
	// gerar a arvore a partir da frequencia dos caracteres do texto
	root := Harvest(GetMap(file, 1))

	// gerar dicionario
	dict := make(map[string]string)
	createDict(root, dict, "")

	//Resetar cursor
	file.Seek(0, 0)
	//Escrever Árvore
	outputFile, _ := os.Create(outputName)
	writer := bit.NewWriter(outputFile)
	writeNode(root, writer)

	// Codificar

	writeCodified(file, dict, writer)
}

//helper
func reverseBits(b byte) byte {
	var d byte
	for i:= 0; i < 8; i++ {
		d <<= 1
		d |= b & 1
		b >>= 1
	}
	return d
}

//Método para ler a arvore recursivamente
func readTree(reader *bit.Reader) *tree.Node{
	read, _ := reader.Read()
	if read { // folha
		char, _ := reader.ReadByte()
		charstring := string(reverseBits(char))
		return tree.New(charstring, nil, nil)
	} else { // tem dois filhos
		leftChild := readTree(reader)
		rightChild := readTree(reader)
		return tree.New("", leftChild, rightChild)
	}
}

func decodeFile(reader *bit.Reader, outputName string, root *tree.Node) {
	output, _ := os.Create(outputName)
	node := root
	for {
		bit, err := reader.Read()
		if err != nil {
			break
		}
		// Anda na árvore, se bit=0 vai pro filho esquerdo
		if bit {
			node = node.Right
		} else {
			node = node.Left
		}

		//Checar se chegamos em uma folha
		if node.IsLeaf() {
			output.WriteString(node.Value)
			node = root
		}
	}
	output.Close()

}
//Recebe um arquivo comprimido (objeto) e retorna o arquivo original (objeto)
func Decompress(file *os.File, outputName string){
	// Ler Árvore (Reconstruir)
	reader := bit.NewReader(file)
	root := readTree(reader)
	if root == nil {
		panic("Árvore nula!")
	}
	// Decodificar percorrendo a arvore

	decodeFile(reader, outputName, root)
}

