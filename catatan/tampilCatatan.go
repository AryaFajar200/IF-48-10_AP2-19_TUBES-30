package catatan

import "fmt"

func TampilkanDaftarCatatan() {
	var i int
	if jumlahCatatan == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}
	fmt.Println("Daftar Catatan:")
	for i = 0; i < jumlahCatatan; i++ {
		fmt.Printf("ID: %d | Tanggal: %s | Topik: %s | Kesulitan: %d\n",
			daftarCatatan[i].id,
			daftarCatatan[i].tanggal,
			daftarCatatan[i].topik,
			daftarCatatan[i].kesulitan)
	}
}
