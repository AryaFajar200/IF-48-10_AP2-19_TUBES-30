package catatan

import "fmt"

func TambahCatatan() {
	if jumlahCatatan >= MAX {
		fmt.Println("Data sudah penuh!")
		return
	}
	fmt.Print("Tanggal (HARI-BULAN-TANGGAL): ")
	fmt.Scan(&daftarCatatan[jumlahCatatan].tanggal)
	fmt.Print("Topik: ")
	fmt.Scan(&daftarCatatan[jumlahCatatan].topik)
	fmt.Print("Isi catatan: ")
	fmt.Scan(&daftarCatatan[jumlahCatatan].isi)
	fmt.Print("Kesulitan (1-5): ")
	fmt.Scan(&daftarCatatan[jumlahCatatan].kesulitan)

	daftarCatatan[jumlahCatatan].id = jumlahCatatan + 1
	jumlahCatatan++
	fmt.Println("Catatan berhasil ditambahkan.")
}
