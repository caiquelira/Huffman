package bitwriter


import "os"

func check(e error) {
	if e != nil {
		panic (e)
	}
}

type BitWriter struct {
	f *os.File
	r int
	bits int
}

func NewBitWriter (file *os.File)(*BitWriter){
	bw := new(BitWriter)
	bw.f = file
	bw.r = 0
	bw.bits = 0
	return bw
}

func (bw *BitWriter) WriteBit (b bool) {
	bw.bits <<= 1
	if (b) {
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

func (bw *BitWriter) WriteByte (b byte) {
	for i := 0; i < 8; i++ {
		bw.WriteBit(bool ((b & 1) == 1))
		b >>= 1
	}
}

func (br *BitWriter) Close (){
	if bw.r == 0 {
		aux := make([]byte, 1)
		aux[0] = byte(0)
		_,err := bw.f.Write(aux)
		check (err)
	} else {
		r := 8 - br.r
		for i := 0; i < r; i++ {
			bw.WriteBit(false)
		}
		aux := make([]byte, 1)
		aux[0] = byte(r)
		_,err := bw.f.Write(aux)
		check (err)
	}
}