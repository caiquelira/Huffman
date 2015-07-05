package huffman

import(
  "priorityQueue"
  "tree"
  "container/heap"
)

// Retorna a arvore
func harvest(freqMap map[string]int) *tree.Node {
  // Primeiro a gente cria a nossa priorityqueue a partir do dicionario de frequencias
  pq := make(priorityQueue, len(freqMap))
  for

}
