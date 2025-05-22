package catatan

import "fmt"

func UbahCatatan(catatan *[MAX]Catatan, jumlah int) {
	var id, i int
	fmt.Print("Masukkan id catatan yang ingin diubah: ")
	fmt.Scan(&id)
	for i = 0; i < JumlahCatatan; i++ {
		if DaftarCatatan[i].ID == id {
			fmt.Print("Tanggal baru: ")
			fmt.Scan(&DaftarCatatan[i].Tanggal)
			fmt.Print("Topik baru: ")
			fmt.Scan(&DaftarCatatan[i].Topik)
			fmt.Print("Isi baru: ")
			fmt.Scan(&DaftarCatatan[i].Isi)
			fmt.Print("Kesulitan baru: ")
			fmt.Scan(&DaftarCatatan[i].Kesulitan)
			fmt.Println("Catatan berhasil diubah (UPDATED YEAY)")
			return
		}
	}
	fmt.Println("Catatan tidak ditemukan.")
}
