package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func BinarySearch(data *catatan.DaftarCatatan, jumlah int, keyword string) {

	var left, right, mid, ketemu int
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
		i := ketemu
		fmt.Println("Topik ditemukan:")
		fmt.Printf("ID        : %d\n", data[i].ID)
		fmt.Printf("Tanggal   : %s\n", data[i].Tanggal)
		fmt.Printf("Topik     : %s\n", data[i].Topik)
		fmt.Printf("Isi       : %s\n", data[i].Isi)
		fmt.Printf("Kesulitan : %d\n", data[i].Kesulitan)
		fmt.Println("----------------------------")
	} else {
		fmt.Println("Topik tidak ditemukan.")
	}
}