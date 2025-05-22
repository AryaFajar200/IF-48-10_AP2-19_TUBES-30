package catatan

import "fmt"

func TambahCatatan(catatan *[MAX]Catatan, jumlah *int) {
	if *jumlah >= MAX {
		fmt.Println("Data penuh!")
		return
	}
	var c Catatan
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scan(&c.Tanggal)
	fmt.Print("Topik: ")
	fmt.Scan(&c.Topik)
	fmt.Print("Isi catatan: ")
	fmt.Scan(&c.Isi)
	fmt.Print("Kesulitan (1-5): ")
	fmt.Scan(&c.Kesulitan)
	c.ID = *jumlah + 1
	catatan[*jumlah] = c
	*jumlah++
	fmt.Println("Catatan berhasil ditambahkan.")
}
