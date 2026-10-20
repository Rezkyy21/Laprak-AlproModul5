package main

import "fmt"

func main() {
	type Buku struct {
		judul        string
		genre        string
		tahunrelease int
		rating       float64
	}

	var buku1, buku2 Buku

	fmt.Println("Masukkan judul buku 1:")
	fmt.Scan(&buku1.judul)
	fmt.Println("Masukkan genre buku 1:")
	fmt.Scan(&buku1.genre)
	fmt.Println("Masukkan tahun rilis buku 1:")
	fmt.Scan(&buku1.tahunrelease)
	fmt.Println("Masukkan rating buku 1:")
	fmt.Scan(&buku1.rating)

	fmt.Println("\nMasukkan judul buku 2:")
	fmt.Scan(&buku2.judul)
	fmt.Println("Masukkan genre buku 2:")
	fmt.Scan(&buku2.genre)
	fmt.Println("Masukkan tahun rilis buku 2:")
	fmt.Scan(&buku2.tahunrelease)
	fmt.Println("Masukkan rating buku 2:")
	fmt.Scan(&buku2.rating)

	fmt.Println("--- Buku 1 ---")
	fmt.Println(buku1.judul)
	fmt.Println(buku1.genre)
	fmt.Println(buku1.tahunrelease)
	fmt.Println(buku1.rating)

	fmt.Println("--- Buku 2 ---")
	fmt.Println(buku2.judul)
	fmt.Println(buku2.genre)
	fmt.Println(buku2.tahunrelease)
	fmt.Println(buku2.rating)
}
