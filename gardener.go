package huffman

import(
  "github.com/caiquelira/huffman/priorityQueue"
  "tree"
  "container/heap"
)

// Retorna a arvore
func harvest(freqMap map[string]int) *tree.Node {
  // Primeiro a gente cria a nossa priorityqueue a partir do dicionario de frequencias
  pq := make(priorityQueue, 0)
  heap.Init(&pq)

  for value, frequency := range freqMap {
    item = &priorityQueue.Item {
           value: value
           frequency := frequency
    }
    heap.Push(&pq, item)
  }

  return nil

}
