package bit

import "strings"

func Example () {
	var ret string
	writer := strings.NewWriter(ret)
	bw := New(writer)
	


	var str string
	str = "Leia essa string"
	reader := strings.NewReader(str)
	br := New(reader, str.size())
}