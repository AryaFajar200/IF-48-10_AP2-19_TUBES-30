package catatan

import "fmt"

func UbahCatatan(catatan *DaftarCatatan, jumlah int) {
	var id, i int
	fmt.Print("Masukkan id catatan yang ingin diubah: ")
	fmt.Scan(&id)
	for i = 0; i < jumlah; i++ {
		if catatan[i].ID == id {
			fmt.Print("Tanggal baru (YYYY-MM-DD): ")
			fmt.Scan(&catatan[i].Tanggal)
			fmt.Print("Topik baru (tidak bisa menggunakan spasi, gunakan underscore): ")
			fmt.Scan(&catatan[i].Topik)
			fmt.Print("Isi baru (tidak bisa menggunakan spasi, gunakan underscore): ")
			fmt.Scan(&catatan[i].Isi)
			fmt.Print("Kesulitan baru: ")
			fmt.Scan(&catatan[i].Kesulitan)
			fmt.Println("Catatan berhasil diubah (UPDATED YEAY)")
			return
		}
	}
	fmt.Println("Catatan tidak ditemukan.")
}
