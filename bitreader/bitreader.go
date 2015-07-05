package huffman


import (
	"os"
	"errors"
	//"bufio"
	//"io"
	//"fmt"
)

func check(e error) {
	if e != nil {
		panic (e)
	}
}

type BitReader struct {
	f *os.File
	max int64
	pos int64
	bits []bool
	r int
}

func (br *BitReader) New (file *os.File)*BitReader{
	br = new(BitReader)
	br.f = file
	br.pos = 0
	br.bits = make([]bool, 8)
	br.r = 0
	file.Seek(-1, 2)
	b := make([]byte, 1)
	_,err := file.Read(b)
	check (err)
	aux,err := file.Stat()
	check(err)
	br.max = aux.Size()*8 - int64(b[0]) - 8
	return br
}

func (br *BitReader) ReadBit ()(b bool, e error) {
	e = nil
	b = false
	if br.pos == br.max {
		e = errors.New("EOF")
		return
	}
	br.pos++
	if br.r > 0 {
		br.r--
		b = br.bits[br.r]
		return
	}
	aux := make([]byte, 1)
	_,err := br.f.Read(aux)
	check(err)
	for i := 0; i < 8; i++ {
		br.bits[i] = bool ((aux[0] & 1) == 1)
		aux[0] >>= 1
	}
	for i := 0; 7 - i > i; i++ {
		br.bits[i], br.bits[7 - i] = br.bits[7 - i], br.bits[i]
	}
	br.r = 8
	br.r--
	b = br.bits[br.r]
	return
}

func (br *BitReader) ReadByte ()(b byte, e error) {
	e = nil
	b = 0
	for i := 0; i < 8; i++ {
		var aux bool
		aux,e = br.ReadBit()
		b <<= 1
		if aux {
			b++
		}
		if e != nil {
			return
		}
	}
	return
}
