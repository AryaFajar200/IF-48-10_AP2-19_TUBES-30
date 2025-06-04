package fitursorting

import (
	"fmt"
	"tubesalpro/catatan"
)

// Selection Sort berdasarkan Tanggal (Ascending)
func SelectionSortByTanggal(data *catatan.DaftarCatatan, jumlah int) {
	var i, j, minIdx int
	for i = 0; i < jumlah-1; i++ {
		minIdx = i
		for j = i + 1; j < jumlah; j++ {
			if data[j].Tanggal < data[minIdx].Tanggal {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tanggal (ascending).")
}
