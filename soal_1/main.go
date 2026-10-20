package main

import "fmt"

type Mahasiswa struct {
	Nama       string
	NIM        string
	Quiz       float64
	Tugas      float64
	Praktikum  float64
	UAS        float64
	UTS        float64
	nilaiAkhir float64
}

const (
	BobotQuiz      = 0.15
	BobotTugas     = 0.15
	BobotPraktikum = 0.20
	BobotUAS       = 0.25
	BobotUTS       = 0.25
)

func main() {
	var m Mahasiswa
	

	fmt.Print("Masukan nama: ")
	fmt.Scanln(&m.Nama)

	fmt.Print("Masukan NIM: ")
	fmt.Scanln(&m.NIM)

	fmt.Print("Masukan nilai Quiz: ")
	fmt.Scanln(&m.Quiz)
	fmt.Print("Masukan nilai Tugas: ")
	fmt.Scanln(&m.Tugas)
	fmt.Print("Masukan nilai Praktikum: ")
	fmt.Scanln(&m.Praktikum)
	fmt.Print("Masukan nilai UAS: ")
	fmt.Scanln(&m.UAS)
	fmt.Print("Masukan nilai UTS: ")
	fmt.Scanln(&m.UTS)

	m.nilaiAkhir = m.Quiz*BobotQuiz +
		m.Tugas*BobotTugas +
		m.Praktikum*BobotPraktikum +
		m.UAS*BobotUAS +
		m.UTS*BobotUTS
	

	fmt.Println()
	fmt.Println("--- Data Mahasiswa ---")
	fmt.Println("NIM ", m.NIM)
	fmt.Println("Nama ", m.Nama)
	fmt.Println("Nilai Quiz ", m.Quiz)
	fmt.Println("Nilai Tugas ", m.Tugas)
	fmt.Println("Nilai Praktikum ", m.Praktikum)
	fmt.Println("Nilai UAS ", m.UAS)
	fmt.Println("Nilai UTS ", m.UTS)
	fmt.Println("Nilai Akhir ", m.nilaiAkhir)
}
