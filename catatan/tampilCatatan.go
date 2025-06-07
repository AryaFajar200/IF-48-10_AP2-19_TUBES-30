package catatan

import "fmt"

func TampilkanDaftarCatatan(catatan *DaftarCatatan, jumlah int) {
	var i int
	if jumlah == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}
	fmt.Println("========================================================================================")
	fmt.Println("|                                 APLIKASI AI ASISTEN BELAJAR                          |")
	fmt.Println("========================================================================================")	
	fmt.Printf("| %-3s | %-20s | %-15s | %-18s | %-3s        |\n", "ID", "TANGGAL", "TOPIK", "ISI", "KESULITAN")
	fmt.Println("========================================================================================")
	for i = 0; i < jumlah; i++ {
	fmt.Printf("| %-3d | %-20s | %-15s | %-18s | %-3d              |\n",
		catatan[i].ID, catatan[i].Tanggal, catatan[i].Topik, catatan[i].Isi, catatan[i].Kesulitan)
	}
	fmt.Println("========================================================================================")

}
