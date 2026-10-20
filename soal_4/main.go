package main

import "fmt"

const (
	diskon1 = 0.10
	diskon2 = 0.20
)

type barang struct {
	nama_barang string
	harga_awal  float64
	harga_akhir float64
}

func main() {
	var b barang

	fmt.Scanln(&b.nama_barang)
	fmt.Scanln(&b.harga_awal)

	b.harga_akhir = b.harga_awal - (b.harga_awal * diskon1)
	b.harga_akhir = b.harga_akhir - (b.harga_akhir * diskon2)

	fmt.Printf("%.2f", b.harga_akhir)
}