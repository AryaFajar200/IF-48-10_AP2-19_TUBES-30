package fitursorting

import (
	"fmt"
	"tubesalpro/catatan"
)

func SelectionSortByTanggal(data *catatan.DaftarCatatan, jumlah int, urutan string) {
	var pass, idx, i int
	var temp catatan.Catatan

	pass = 1
	for pass <= jumlah-1 {
		idx = pass - 1
		i = pass
		for i < jumlah {
			if (urutan == "ascending" && data[idx].Tanggal > data[i].Tanggal) ||
				(urutan == "descending" && data[idx].Tanggal < data[i].Tanggal) {
				idx = i
			}
			i++
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
		pass++
	}

	if urutan == "ascending" {
		fmt.Println("Data berhasil diurutkan berdasarkan tanggal (ascending).")
	} else if urutan == "descending" {
		fmt.Println("Data berhasil diurutkan berdasarkan tanggal (descending).")
	} else {
		fmt.Println("Pilihan urutan tidak valid (gunakan 'ascending' atau 'descending').")
	}
}
