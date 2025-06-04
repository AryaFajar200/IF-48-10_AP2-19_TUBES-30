package catatan

import "fmt"

func UbahCatatan(catatan *DaftarCatatan, jumlah int) {
	var id, i int
	fmt.Print("Masukkan id catatan yang ingin diubah: ")
	fmt.Scan(&id)
	for i = 0; i < jumlah; i++ {
		if catatan[i].ID == id {
			fmt.Print("Tanggal baru: ")
			fmt.Scan(&catatan[i].Tanggal)
			fmt.Print("Topik baru: ")
			fmt.Scan(&catatan[i].Topik)
			fmt.Print("Isi baru: ")
			fmt.Scan(&catatan[i].Isi)
			fmt.Print("Kesulitan baru: ")
			fmt.Scan(&catatan[i].Kesulitan)
			fmt.Println("Catatan berhasil diubah (UPDATED YEAY)")
			return
		}
	}
	fmt.Println("Catatan tidak ditemukan.")
}
