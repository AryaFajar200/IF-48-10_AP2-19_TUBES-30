package catatan

type Catatan struct {
	id        int
	tanggal string
	topik     string
	isi       string
	kesulitan int
}

const MAX = 100

var DaftarCatatan [MAX]Catatan
var JumlahCatatan int = 0
