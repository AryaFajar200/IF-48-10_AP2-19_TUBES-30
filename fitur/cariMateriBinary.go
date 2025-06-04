package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func BinarySearch(data *catatan.DaftarCatatan, jumlah int, keyword string) {
	var low, high, mid int
    SortByTopik(data, jumlah)

    low = 0
    high = jumlah - 1

    for low <= high {
        mid = (low + high) / 2
        if data[mid].Topik == keyword {
            fmt.Println("Topik ditemukan:")
			fmt.Printf("ID        : %d\n", data[mid].ID)
			fmt.Printf("Tanggal   : %s\n", data[mid].Tanggal)
			fmt.Printf("Topik     : %s\n", data[mid].Topik)
			fmt.Printf("Isi       : %s\n", data[mid].Isi)
			fmt.Printf("Kesulitan : %d\n", data[mid].Kesulitan)
			fmt.Println("----------------------------")
            return
        } else if data[mid].Topik < keyword {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    fmt.Println("Topik tidak ditemukan.")
}