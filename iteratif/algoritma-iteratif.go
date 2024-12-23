package main

import (
	"fmt"
	"time"
)

type Kunjungan struct {
	Tanggal  string
	Jam      string
	Kategori string
	Tujuan   string
	Durasi   int
}

var dataset = []Kunjungan{
	{"2024-12-14", "09:00", "Mahasiswa", "Membaca", 2},
	{"2024-12-14", "10:00", "Dosen", "Mengembalikan Buku", 1},
	{"2024-12-15", "11:00", "Mahasiswa", "Meminjam Buku", 3},
	{"2024-12-16", "08:00", "Umum", "Membaca", 2},
	{"2024-12-16", "14:00", "Mahasiswa", "Menggunakan Komputer", 4},
	{"2024-12-17", "09:00", "Dosen", "Membaca", 1},
	{"2024-12-17", "13:00", "Mahasiswa", "Membaca", 2},
	{"2024-12-17", "15:00", "Umum", "Menggunakan Komputer", 2},
	{"2024-12-18", "10:00", "Mahasiswa", "Meminjam Buku", 3},
	{"2024-12-18", "11:00", "Dosen", "Mengembalikan Buku", 1},
	{"2024-12-19", "09:00", "Umum", "Membaca", 2},
	{"2024-12-19", "13:00", "Mahasiswa", "Menggunakan Komputer", 4},
	{"2024-12-20", "10:00", "Mahasiswa", "Meminjam Buku", 3},
	{"2024-12-20", "11:00", "Dosen", "Membaca", 1},
	{"2024-12-20", "15:00", "Umum", "Menggunakan Komputer", 2},
}

// FUNGSI MENGHITUNG JUMLAH KUNJUNGAN BERDASARKAN KATEGORI
func hitungIteratif(data []Kunjungan, kategori string) int {
	count := 0
	for _, kunjungan := range data {
		if kunjungan.Kategori == kategori {
			count++
		}
	}
	return count
}

// FUNGSI MENGHITUNG KUNJUNGAN PERHARI
func hitungKunjunganPerHari(data []Kunjungan) map[string]int {
	kunjunganPerHari := make(map[string]int)
	for _, kunjungan := range data {
		kunjunganPerHari[kunjungan.Tanggal]++
	}
	return kunjunganPerHari
}

// FUNGSI MENGHITUNG TUJUAN KUNJUNGAN PALING SERING
func hitungTujuanPalingSering(data []Kunjungan) map[string]int {
	tujuanSering := make(map[string]int)
	for _, kunjungan := range data {
		tujuanSering[kunjungan.Tujuan]++
	}
	return tujuanSering
}

// FUNGSI MENGHITUNG JAM KUNJUNGAN PALING SERING
func hitungJamPalingSering(data []Kunjungan) map[string]int {
	jamSering := make(map[string]int)
	for _, kunjungan := range data {
		jamSering[kunjungan.Jam]++
	}
	return jamSering
}

func main() {
	start := time.Now()

	// HITUNG TOTAL KUNJUNGAN
	totalMahasiswa := hitungIteratif(dataset, "Mahasiswa")
	totalDosen := hitungIteratif(dataset, "Dosen")
	totalUmum := hitungIteratif(dataset, "Umum")

	// HITUNG JUMLAH KUNJUNGAN PERHARI
	kunjunganPerHari := hitungKunjunganPerHari(dataset)

	// HITUNG TUJUAN KUNJUNGAN PALING SERING
	tujuanSering := hitungTujuanPalingSering(dataset)

	// HITUNG JAM KUNJUNGAN PALING SERING
	jamSering := hitungJamPalingSering(dataset)

	fmt.Println("=== ALGORITMA ITERATIF === ")
	fmt.Println("= JUMLAH KUNJUGAN")
	fmt.Printf("Mahasiswa : %d, Dosen : %d, Umum : %d\n", totalMahasiswa, totalDosen, totalUmum)

	fmt.Println("\n= JUMLAH KUNJUNGAN PERHARI")

	for tanggal, jumlah := range kunjunganPerHari {
		fmt.Printf("%s : %d\n", tanggal, jumlah)
	}

	fmt.Println("\n= TUJUAN KUNJUNGAN")
	for tujuan, jumlah := range tujuanSering {
		fmt.Printf("%s : %d\n", tujuan, jumlah)
	}

	fmt.Println("\n= JAM KUNJUNGAN")
	for jam, jumlah := range jamSering {
		fmt.Printf("%s : %d\n", jam, jumlah)
	}

	duration := time.Since(start).Seconds()
	fmt.Printf("\nWaktu Eksekusi : %.6f detik\n", duration)
}
