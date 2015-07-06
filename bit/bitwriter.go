package bit


import (
	"os"
	"io"
)

func check(e error) {
	if e != nil {
		panic (e)
	}
}

type writer struct {
	f io.Writer
	r int
	bits int
}

type Writer struct {
	f *os.File
	b *writer
}

func newWriter(file io.Writer)(*writer){
	bw := new(writer)
	bw.f = file
	bw.r = 0
	bw.bits = 0
	return bw
}

func NewWriter(file *os.File)(*Writer){
	bw := new(Writer)
	bw.b = newWriter(file)
	return bw
}

func (bw *writer) write(in bool) {
	bw.bits <<= 1
	if (in) {
		bw.bits++
	}
	bw.r++
	if (bw.r == 8) {
		bw.r = 0
		x := make ([]byte, 1)
		x[0] = byte(bw.bits)
		_, err := bw.f.Write(x)
		check (err)
		bw.bits = 0
	}
}

func (bw *Writer) Write(in bool) {
	bw.b.write(in)
}

func (bw *Writer) WriteByte (b byte) {
	for i := 0; i < 8; i++ {
		bw.Write(bool ((b & 1) == 1))
		b >>= 1
	}
}

func (bw *writer) close (){
	if bw.r > 0 {
		r := 8 - bw.r
		for i := 0; i < r; i++ {
			bw.write(false)
		}
		aux := make([]byte, 1)
		aux[0] = byte(r)
		_,err := bw.f.Write(aux)
		check (err)
	}
}

func (bw *Writer) Close(){
	bw.b.close()
}