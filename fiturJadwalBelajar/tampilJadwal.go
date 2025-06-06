package fiturJadwalBelajar

import (
	"fmt"
	"tubesalpro/catatan"
)

func TampilkanJadwalBelajar(data *catatan.DaftarCatatan, jumlah int) {
	var hari, i, j int
	if jumlah == 0 {
		fmt.Println("Belum ada catatan untuk dijadwalkan.")
		return
	}

	var slotPerHari int
	fmt.Print("Berapa slot waktu luang Anda per hari (contoh: 2)? ")
	fmt.Scan(&slotPerHari)

	if slotPerHari <= 0 {
		fmt.Println("Slot per hari tidak valid.")
		return
	}

	hari = 1
	fmt.Printf("\nJadwal Belajar:\n")
	for i = 0; i < jumlah; i += slotPerHari {
		fmt.Printf("\nHari ke-%d:\n", hari)
		for j = 0; j < slotPerHari && i+j < jumlah; j++ {
			fmt.Printf("- [%s] %s\n", data[i+j].Tanggal, data[i+j].Topik)
		}
		hari++
	}
}
