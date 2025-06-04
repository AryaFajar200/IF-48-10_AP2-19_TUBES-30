package catatan

type Catatan struct {
	ID      int
	Tanggal string
	Topik     string
	Isi       string
	Kesulitan int
}

type DaftarCatatan [100]Catatan

const MAX = 100

