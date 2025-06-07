package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func SequentialSearch(data *catatan.DaftarCatatan, jumlah int, keyword string) {
	var ketemu int = -1
	var k int = 0

	for ketemu == -1 && k < jumlah {
		if data[k].Topik == keyword {
			ketemu = k
		}
		k++
	}
	var i int

	if ketemu != -1 {
		i = ketemu
		fmt.Println("+----------------------------+")
		fmt.Println("|      Topik ditemukan:      |")
		fmt.Println("+----------------------------+")
		fmt.Printf("| %-12s : %-10d |\n", "ID", data[i].ID)
		fmt.Printf("| %-12s : %-10s |\n", "Tanggal", data[i].Tanggal)
		fmt.Printf("| %-12s : %-10s |\n", "Topik", data[i].Topik)
		fmt.Printf("| %-12s : %-10s |\n", "Isi", data[i].Isi)
		fmt.Printf("| %-12s : %-10d |\n", "Kesulitan", data[i].Kesulitan)
		fmt.Println("+----------------------------+")
	} else {
		fmt.Printf("Topik dengan keyword %s tidak ditemukan.\n", keyword)
	}
}



