package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func InsertionSortByKesulitan(data *catatan.DaftarCatatan, jumlah int, urutan string) {
	var pass, i int
	var temp catatan.Catatan

	pass = 1
	for pass <= jumlah-1 {
		i = pass
		temp = data[pass]

		if urutan == "ascending" {
			for i > 0 && temp.Kesulitan < data[i-1].Kesulitan {
				data[i] = data[i-1]
				i--
			}
		} else if urutan == "descending" {
			for i > 0 && temp.Kesulitan > data[i-1].Kesulitan {
				data[i] = data[i-1]
				i--
			}
		} else {
			fmt.Println("Pilihan urutan tidak valid (gunakan 'ascending' atau 'descending').")
			return
		}

		data[i] = temp
		pass++
	}

	if urutan == "ascending" {
		fmt.Println("Data berhasil diurutkan berdasarkan kesulitan (ascending).")
	} else {
		fmt.Println("Data berhasil diurutkan berdasarkan kesulitan (descending).")
	}
}
