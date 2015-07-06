package bit


import (
	"os"
	"io"
	"errors"
	//"bufio"
	//"fmt"
)

type reader struct {
	f io.Reader
	max int64
	pos int64
	bits []bool
	r int
}

type Reader struct {
	f *os.File
	b *reader
}

func newReader(file io.ReadSeeker, max int64)*reader{
	br := new(reader)
	br.f = file
	br.pos = 0
	br.bits = make([]bool, 8)
	br.r = 0
	br.max = max
	return br
}

func NewReader(file *os.File)*Reader{
	br := new(Reader)
	br.f = file;
	file.Seek(-1, 2)
	b := make([]byte, 1)
	_,err := file.Read(b)
	file.Seek(0,0)
	check (err)
	aux,err := file.Stat()
	check(err)
	br.b = newReader(file, aux.Size()*8 - int64(b[0]) - 8)
	return br
}

func (br *reader) read()(b bool, e error) {
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
	br.r = 8
	br.r--
	b = br.bits[br.r]
	return
}

func (br *Reader) Read()(b bool, e error) {
	return br.b.read()
}
func reverseBits(b byte) byte {
	var d byte
	for b > 0 {
		d |= b & 1
		b >>= 1
		d <<= 1
	}
	return d
}

func (br *Reader) ReadByte()(b byte, e error) {
	e = nil
	b = 0
	for i := 0; i < 8; i++ {
		var aux bool
		aux,e = br.Read()
		b <<= 1
		if aux {
			b++
		}
		if e != nil {
			reverseBits(b)
			return
		}
	}
	reverseBits(b)
	return
}
