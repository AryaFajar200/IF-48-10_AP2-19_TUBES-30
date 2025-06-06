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



