package PriorityQueue

import(
  "tree"
)

type Item struct {
  node *tree.Node // Cada elemento da arvore
  frequency int // Quantas vezes o char apareceu no texto

  // O index eh necessario para atualizar a heap e eh mantido pelos metodos da interface
  index int
}

// Vamos fazer uma fila de prioridade para pegar os elementos a serem adicionados na arvore
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
  return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
  // Queremos que pop retorne o elemento de maior prioridade
  return pq[i].frequency > pq[j].frequency
}

// Nome auto explicativo, do ingles trocar
func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i
  pq[j].index = j
}

// Para implementarmos a interface heap temos que ter Push() e Pop(), com os parametros como os abaixo
func (pq *PriorityQueue) Push(x interface{}) {
  n := len(*pq)
  item := x.(*Item)
  item.index = n
  *pq = append(*pq, item)
}

// Retorna o elemento com a maior prioridade ( primeiro da fila do Less())
func (pq *PriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  item := old[n-1]
  item.index = -1 // Para nao acessarmos nada existente
  *pq = old[0 : n-1]
  return item
}
