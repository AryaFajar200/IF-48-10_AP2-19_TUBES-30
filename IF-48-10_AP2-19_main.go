package main

import (
	"fmt"
	"tubesalpro/catatan"
	"tubesalpro/fitur"
)

func main() {
	var pilihan, nData int
	var data catatan.DaftarCatatan
	for {
		fmt.Println("=============================================================")
		fmt.Println("|                 APLIKASI AI ASISTEN BELAJAR               |")
		fmt.Println("=============================================================")
		fmt.Println("| 1. Tambah Catatan                                         |")
		fmt.Println("| 2. Ubah Catatan                                           |")
		fmt.Println("| 3. Hapus Catatan                                          |")
		fmt.Println("| 4. Tampilkan Catatan                                      |")
		fmt.Println("| 5. Cari Catatan (Sequential Search)                       |")
		fmt.Println("| 6. Cari Catatan (Binary Search)                           |")
		fmt.Println("| 7. Urutkan Catatan Berdasarkan Tanggal (Selection Sort)   |")
		fmt.Println("| 8. Urutkan Catatan Berdasarkan Kesulitan (Insertion Sort) |")
		fmt.Println("| 9. Buat Jadwal Belajar                                    |")
		fmt.Println("| 10. Soal Latihan (interaktif)                             |")
		fmt.Println("| 0. Keluar                                                 |")
		fmt.Println("=============================================================")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			catatan.TambahCatatan(&data, &nData)
		case 2:
			catatan.UbahCatatan(&data, nData)
		case 3:
			catatan.HapusCatatan(&data, &nData)
		case 4:
			catatan.TampilkanDaftarCatatan(&data, nData)
		case 5:
   			 var keyword string
   			 fmt.Print("Masukkan keyword topik: ")
    		 fmt.Scan(&keyword)
   			 fitur.SequentialSearch(&data, nData, keyword)
		case 6:
   			 var keyword string
   			 fmt.Print("Masukkan keyword topik: ")
   			 fmt.Scan(&keyword)
   			 fitur.BinarySearch(&data, nData, keyword)
		case 7:
			var opsi string
			fmt.Print("Urutkan berdasarkan tanggal [ascending/descending]: ")
			fmt.Scan(&opsi)
			fitur.SelectionSortByTanggal(&data, nData, opsi)
			catatan.TampilkanDaftarCatatan(&data, nData)
		case 8:
			var opsi string
			fmt.Print("Urutkan berdasarkan kesulitan [ascending/descending]: ")
			fmt.Scan(&opsi)
			fitur.InsertionSortByKesulitan(&data, nData, opsi)
			catatan.TampilkanDaftarCatatan(&data, nData)
		case 9:
			fitur.TampilkanJadwalBelajar(&data, nData)
		case 10:
			fitur.TampilkanSoalLatihan(&data, nData)
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
