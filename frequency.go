package huffman

import (
	"io"
	"bufio"
	//"fmt"
)

func GetMap(fi io.Reader) map[string]int{
	file := bufio.NewReader(fi)
	//Dicionario que relaciona cada string com sua frequencia
	freqMap := make(map[string]int)
	// Buffer para lermos 1 rune de cada vez
	//
	for {
		// n eh o tamanho do array de bytes que o Reader retorna
		r, _, err := file.ReadRune()
		if err != nil && err != io.EOF {
			panic(err) // Se houver um erro diferente do arquivo ter acabado
		}

		if err == io.EOF {
			break
		}

		// Transforma a leitura do Reader numa string
		str := string(r)
		// Aumenta a frequencia do elemento com o valor da string
		// ou adiciona um novo key do Dicionario
		//fmt.Println("Frequency of char:", str, "\trune:", r," \tbyte:", byte(r), "\tincreased.")
		freqMap[str]++
	}

	return freqMap
}
