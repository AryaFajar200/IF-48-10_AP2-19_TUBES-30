package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func BinarySearch(data *catatan.DaftarCatatan, jumlah int, keyword string) {

	var left, right, mid, ketemu, i int
	left = 0
	right = jumlah - 1
	ketemu = -1

	for left <= right && ketemu == -1 {
		mid = (left + right) / 2
		if data[mid].Topik == keyword {
			ketemu = mid
		} else if data[mid].Topik < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

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
}else {
		fmt.Printf("Topik dengan keyword %s tidak ditemukan.\n", keyword)
	}
}