package fitur

import (
	"fmt"
	"tubesalpro/catatan"
)

func TampilkanSoalLatihan(data *catatan.DaftarCatatan, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada catatan.")
		return
	}

	var i, pilih int
	var topik, isi, jawaban string
	var benar bool

	fmt.Println("\n+--------------------------+")
	fmt.Println("|     Daftar Topik        |")
	fmt.Println("+--------------------------+")
	for i = 0; i < jumlah; i++ {
		fmt.Printf("| %-2d. %-21s |\n", i+1, data[i].Topik)
	}
	fmt.Println("+--------------------------+")

	fmt.Print("\nPilih nomor topik yang ingin dijadikan soal: ")
	fmt.Scan(&pilih)

	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	topik = data[pilih-1].Topik
	isi = data[pilih-1].Isi

	fmt.Println("\n+---------------------------------------------+")
	fmt.Printf("| Soal: Apa yang dimaksud dengan \"%s\"? |\n", topik)
	fmt.Println("+---------------------------------------------+")

	fmt.Print("Jawaban Anda (1 kata / tanpa spasi): ")
	fmt.Scan(&jawaban)

	if jawaban == isi {
		benar = true
	} else {
		benar = false
	}

	fmt.Println("+--------------------------+")
	fmt.Printf("| Jawaban benar? %-5v       |\n", benar)
	fmt.Println("+--------------------------+")
}
