package huffman

import(
  "github.com/caiquelira/huffman/huffmanHeap"
  "github.com/caiquelira/huffman/tree"
  "container/heap"
  "errors"
)

// Retorna a arvore
func harvest(freqMap map[string]int) *tree.Node {
  // Primeiro a gente cria a nossa priorityqueue a partir do dicionario de frequencias
  hh := huffmanHeap.New(freqMap)

  for {
    // Caso a heap contenha um unico elemento retornamos ela
    if hh.Len() == 1 {
      return heap.Pop(&hh).(*huffmanHeap.Item).Node
    }

    // Caso contrario pegamos os dois elementos do topo
    r, ok1 := heap.Pop(&hh).(*huffmanHeap.Item)
    l, ok2 := heap.Pop(&hh).(*huffmanHeap.Item)
    if !ok1 || !ok2 {
      panic(errors.New("Element was not of type huffmanHeap.Item"))
    }

    //E criamos uma "arvore" intermediaria com eles
    newItem := &huffmanHeap.Item{
               Node: tree.New("", r.Node, l.Node),
               Frequency: r.Frequency + l.Frequency,
    }
    // Adicionando em seguida ao
    heap.Push(&hh, newItem)
  }
}
