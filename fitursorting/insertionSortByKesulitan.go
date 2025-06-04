package fitursorting

import (
	"fmt"
	"tubesalpro/catatan"
)

// Insertion Sort berdasarkan Kesulitan (Ascending)
func InsertionSortByKesulitan(data *catatan.DaftarCatatan, jumlah int) {
	var i, j int
	var temp catatan.Catatan
	for i = 1; i < jumlah; i++ {
		temp = data[i]
		j = i - 1
		for j >= 0 && data[j].Kesulitan > temp.Kesulitan {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan kesulitan (ascending).")
}

