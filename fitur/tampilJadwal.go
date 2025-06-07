package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func TampilkanJadwalBelajar(data *catatan.DaftarCatatan, jumlah int) {
	var hari, i, j, slotPerHari int

	if jumlah == 0 {
		fmt.Println("Belum ada catatan untuk dijadwalkan.")
		return
	}

	SelectionSortByTanggal(data, jumlah, "ascending")

	fmt.Print("Berapa slot waktu luang Anda per hari (contoh: 2)? ")
	fmt.Scan(&slotPerHari)

	if slotPerHari <= 0 {
		fmt.Println("Slot per hari tidak valid.")
		return
	}

	hari = 1
	fmt.Println("\n+=========================================+")
	fmt.Println("|             JADWAL BELAJAR             |")
	fmt.Println("+=========================================+")

	for i = 0; i < jumlah; i += slotPerHari {
		fmt.Printf("\n+----------------------+\n")
		fmt.Printf("| Hari ke-%-13d |\n", hari)
		fmt.Println("+----------------------+")
		fmt.Printf("| %-10s | %-10s |\n", "Tanggal", "Topik")
		fmt.Println("+------------+------------+")

		for j = 0; j < slotPerHari && i+j < jumlah; j++ {
			fmt.Printf("| %-10s | %-10s |\n", data[i+j].Tanggal, data[i+j].Topik)
		}

		fmt.Println("+------------+------------+")
		hari++
	}
}
