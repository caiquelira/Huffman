package huffman

import (
	"fmt"
	"github.com/caiquelira/huffman/tree"
)

func ExampleGardener() {
	freqMap := map[string]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4,
	}

	t := harvest(freqMap)

	fmt.Print(t)

	// Output:
	// ""
	//   "d"
	//   ""
	//     "c"
	//     ""
	//       "a"
	//       "b"

}
