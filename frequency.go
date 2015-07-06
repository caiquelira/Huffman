package huffman

import (
	"io"
)

func GetMap(file io.Reader, readerSize int) map[string]int{
	//Dicionario que relaciona cada string com sua frequencia
	freqMap := make(map[string]int)
	// Buffer para lermos 1 byte de cada vez
	buf := make([]byte, readerSize)
	//
	for {
		// n eh o tamanho do array de bytes que o Reader retorna
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err) // Se houver um erro diferente do arquivo ter acabado
		}
		if n == 0 {
			break // Paramos de ler quando o tamanho do buffer for nulo
		}
		// Transforma a leitura do Reader numa string
		str := string(buf[:n])
		// Aumenta a frequencia do elemento com o valor da string
		// ou adiciona um novo key do Dicionario
		freqMap[str]++
	}
	return freqMap
}
