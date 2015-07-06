package huffman

import (
	"fmt"
)

func ExampleGardener() {
	freqMap := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4,
	}

	t := Harvest(freqMap)

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
