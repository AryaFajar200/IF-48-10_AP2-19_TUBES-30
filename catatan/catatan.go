package catatan

type Catatan struct {
	ID      int
	Tanggal string
	Topik     string
	Isi       string
	Kesulitan int
}

const MAX = 100

var DaftarCatatan [MAX]Catatan
var JumlahCatatan int = 0
