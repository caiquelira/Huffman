package huffman

import (
	"fmt"
	"strings"
)

func ExampleFrequency() {
	testString := "abbcccdddd"
	reader := strings.NewReader(testString)
	freqMap := getMap(reader, 1)
	for k, v := range freqMap {
		fmt.Println("%s:%s", k, v)
	}

	// Output:
	// a:1
	// b:2
	// c:3
	// d:4
}
