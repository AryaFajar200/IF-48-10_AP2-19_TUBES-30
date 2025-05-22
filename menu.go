package main

import (
	"fmt"
	"tubesalpro/catatan"
	"tubesalpro/fitur"
)



func menu() {
	var pilihan int
	for {
		fmt.Println("\n=== APLIKASI AI ASISTEN BELAJAR ===")
		fmt.Println("1. Tambah Catatan")
		fmt.Println("2. Ubah Catatan")
		fmt.Println("3. Hapus Catatan")
		fmt.Println("4. Tampilkan Catatan")
		fmt.Println("5. Cari Catatan (Sequential Search)")
		fmt.Println("6. Cari Catatan (Binary Search)")
		// fmt.Println("6. Urutkan Catatan Berdasarkan Tanggal (Selection Sort)")
		// fmt.Println("7. Urutkan Catatan Berdasarkan Kesulitan (Insertion Sort)")
		// fmt.Println("8. Tampilkan Soal Latihan")
		// fmt.Println("9. Buat Jadwal Belajar")
		fmt.Println("0. Keluar")
		fmt.Println("===================================")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			catatan.TambahCatatan(&catatan.DaftarCatatan, &catatan.JumlahCatatan)
		case 2:
			catatan.UbahCatatan(&catatan.DaftarCatatan, catatan.JumlahCatatan)
		case 3:
			catatan.HapusCatatan(&catatan.DaftarCatatan, &catatan.JumlahCatatan)
		case 4:
			catatan.TampilkanDaftarCatatan(&catatan.DaftarCatatan, catatan.JumlahCatatan)
		case 5:
   			 var keyword string
   			 fmt.Print("Masukkan keyword topik: ")
    		 fmt.Scan(&keyword)
   			 fitur.SequentialSearch(&catatan.DaftarCatatan, catatan.JumlahCatatan, keyword)
		// case 6:
   		// 	 var keyword string
   		// 	 fmt.Print("Masukkan keyword topik: ")
   		// 	 fmt.Scan(&keyword)
   		// 	 fitur.BinarySearch(&catatan.DaftarCatatan, catatan.JumlahCatatan, keyword)
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
