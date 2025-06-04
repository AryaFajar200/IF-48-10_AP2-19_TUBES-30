package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func SequentialSearch(data *catatan.DaftarCatatan, jumlah int, keyword string) {
	var ketemu bool
	var i int
	ketemu = false
	for i = 0; i < jumlah; i++ {
		if data[i].Topik == keyword {
			fmt.Println("Topik ditemukan:")
			fmt.Printf("ID        : %d\n", data[i].ID)
			fmt.Printf("Tanggal   : %s\n", data[i].Tanggal)
			fmt.Printf("Topik     : %s\n", data[i].Topik)
			fmt.Printf("Isi       : %s\n", data[i].Isi)
			fmt.Printf("Kesulitan : %d\n", data[i].Kesulitan)
			fmt.Println("----------------------------")
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Topik tidak ditemukan.")
	}
}


func SortByTopik(data *catatan.DaftarCatatan, jumlah int) {
	var i, j int
    for i = 0; i < jumlah-1; i++ {
        for j = 0; j < jumlah-i-1; j++ {
            if data[j].Topik > data[j+1].Topik {
                data[j], data[j+1] = data[j+1], data[j]
            }
        }
    }
}
