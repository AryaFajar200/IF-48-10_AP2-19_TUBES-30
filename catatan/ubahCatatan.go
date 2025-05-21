package catatan

import "fmt"

func UbahCatatan() {
	var id, i int
	fmt.Print("Masukkan id catatan yang ingin diubah: ")
	fmt.Scan(&id)
	for i = 0; i < jumlahCatatan; i++ {
		if daftarCatatan[i].id == id {
			fmt.Print("Tanggal baru: ")
			fmt.Scan(&daftarCatatan[i].tanggal)
			fmt.Print("Topik baru: ")
			fmt.Scan(&daftarCatatan[i].topik)
			fmt.Print("Isi baru: ")
			fmt.Scan(&daftarCatatan[i].isi)
			fmt.Print("Kesulitan baru: ")
			fmt.Scan(&daftarCatatan[i].kesulitan)
			fmt.Println("Catatan berhasil diubah (UPDATED YEAY)")
			return
		}
	}
	fmt.Println("Catatan tidak ditemukan.")
}
