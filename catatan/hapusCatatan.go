package catatan

import "fmt"

func HapusCatatan() {
	var id, i, j int
	fmt.Print("Masukkan id catatan yang ingin dihapus: ")
	fmt.Scan(&id)
	for i = 0; i < jumlahCatatan; i++ {
		if daftarCatatan[i].id == id {
			for j = i; j < jumlahCatatan-1; j++ {
				daftarCatatan[j] = daftarCatatan[j+1]
			}
			jumlahCatatan--
			fmt.Println("Catatan berhasil dihapus.")
			return
		}
	}
	fmt.Println("Catatan tidak ditemukan.")
}
