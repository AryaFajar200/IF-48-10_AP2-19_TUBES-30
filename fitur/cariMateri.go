package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func SequentialSearch(data *[catatan.MAX]catatan.Catatan, jumlah int, keyword string) {
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

// func BinarySearch(data *[catatan.MAX]catatan.Catatan, jumlah int, keyword string) {
// 	var low, high, mid int
//     SortByTopik(data, jumlah)

//     low = 0
//     high = jumlah - 1

//     for low <= high {
//         mid = (low + high) / 2
//         if data[mid].Topik == keyword {
//             fmt.Println("Ditemukan:", data[mid])
//             return
//         } else if data[mid].Topik < keyword {
//             low = mid + 1
//         } else {
//             high = mid - 1
//         }
//     }
//     fmt.Println("Topik tidak ditemukan.")
// }


func SortByTopik(data *[catatan.MAX]catatan.Catatan, jumlah int) {
	var i, j int
    for i = 0; i < jumlah-1; i++ {
        for j = 0; j < jumlah-i-1; j++ {
            if data[j].Topik > data[j+1].Topik {
                data[j], data[j+1] = data[j+1], data[j]
            }
        }
    }
}
