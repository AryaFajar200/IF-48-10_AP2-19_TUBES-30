package catatan

import "fmt"

func HapusCatatan(catatan *[MAX]Catatan, jumlah *int) {
	var id, i, j int
	fmt.Print("Masukkan ID catatan yang ingin dihapus: ")
	fmt.Scan(&id)
	for i = 0; i < *jumlah; i++ {
		if catatan[i].id == id {
			for j = i; j < *jumlah-1; j++ {
				catatan[j] = catatan[j+1]
			}
			*jumlah--
			fmt.Println("Catatan berhasil dihapus.")
			return
		}
	}
	fmt.Println("Catatan tidak ditemukan.")
}
