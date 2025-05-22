package catatan

import "fmt"

func TampilkanDaftarCatatan(catatan *[MAX]Catatan, jumlah int) {
	var i int
	if jumlah == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}
	fmt.Println("Daftar Catatan:")
	for i = 0; i < jumlah; i++ {
		fmt.Printf("ID: %d | Tanggal: %s | Topik: %s | Kesulitan: %d\n",
			catatan[i].ID, catatan[i].Tanggal, catatan[i].Topik, catatan[i].Kesulitan)
	}
}
