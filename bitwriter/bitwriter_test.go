package bitwriter

import "strings"

func Example () {
	var ret string
	writer := strings.NewWriter(ret)
	bw := New(writer)
	
}