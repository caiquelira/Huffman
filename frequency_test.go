package huffman

import (
	"reflect"
	"strings"
	"testing"
)

func TestFrequency(t *testing.T) {
	testString := "abbcccddddé"
	expectedMap := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "é": 1,
	}

	reader := strings.NewReader(testString)
	freqMap := GetMap(reader)

	eq := reflect.DeepEqual(freqMap, expectedMap)
	if !eq {
		t.Error(freqMap, "different than expected:", expectedMap)
	}
}
