package huffmanHeap

import(
  "container/heap"
  "tree"
)

type Item struct {
  node *tree.Node // Cada elemento da arvore
  frequency int // Quantas vezes o char apareceu no texto

  // O index eh necessario para atualizar a heap e eh mantido pelos metodos da interface
  index int
}

// "Construtor" da huffmanHeap
func New(map[string]int) huffmanHeap{
  hh := make(huffmanHeap, 0)
  heap.Init(&hh)

  //
  for value, frequency := range freqMap {
    item = &huffmanHeap.Item {
           node: tree
           frequency := frequency
    }
    heap.Push(&hh, item)
  }

  return hh
}

// Vamos fazer uma fila de prioridade para pegar os elementos a serem adicionados na arvore
type huffmanHeap []*Item

func (hh huffmanHeap) Len() int {
  return len(hh)
}

func (hh huffmanHeap) Less(i, j int) bool {
  // Queremos que pop retorne o elemento de maior prioridade
  return hh[i].frequency > hh[j].frequency
}

// Nome auto explicativo, do ingles trocar
func (hh huffmanHeap) Swap(i, j int) {
  hh[i], hh[j] = hh[j], hh[i]
  hh[i].index = i
  hh[j].index = j
}

// Para implementarmos a interface heap temos que ter Push() e Pop(), com os parametros como os abaixo
func (hh *huffmanHeap) Push(x interface{}) {
  n := len(*hh)
  item := x.(*Item)
  item.index = n
  *hh = append(*hh, item)
}

// Retorna o elemento com a maior prioridade ( primeiro da fila do Less())
func (hh *huffmanHeap) Pop() interface{} {
  old := *hh
  n := len(old)
  item := old[n-1]
  item.index = -1 // Para nao acessarmos nada existente
  *hh = old[0 : n-1]
  return item
}
