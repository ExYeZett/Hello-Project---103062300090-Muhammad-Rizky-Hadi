package main

import (
	"fmt"
)

type mahasiswa struct {
	No         int
	NIM        string
	Nama       string
	IPK        float64
}

func urutkanGPA(daftarMahasiswa []mahasiswa) {
	n := len(daftarMahasiswa)
	for pass := 0; pass < n-1; pass++ {
		minIdx := pass
		for i := pass + 1; i < n; i++ {
			if daftarMahasiswa[i].IPK < daftarMahasiswa[minIdx].IPK {
				minIdx = i
			}
		}
		daftarMahasiswa[pass], daftarMahasiswa[minIdx] = daftarMahasiswa[minIdx], daftarMahasiswa[pass]
	}
}

func main() {
	daftarMahasiswa := []mahasiswa{
		{1, "113211000", "Andi", 3.10},
		{2, "113211001", "Budi", 2.90},
		{3, "113211002", "Citra", 3.25},
		{4, "113211003", "Dewi", 2.75},
	}

	fmt.Println("Sebelum diurutkan:")
	for _, m := range daftarMahasiswa {
		fmt.Printf("No: %d, NIM: %s, Nama: %s, IPK: %.2f\n", m.No, m.NIM, m.Nama, m.IPK)
	}

	urutkanGPA(daftarMahasiswa)

	fmt.Println("\nSetelah diurutkan berdasarkan IPK:")
	for _, m := range daftarMahasiswa {
		fmt.Printf("No: %d, NIM: %s, Nama: %s, IPK: %.2f\n", m.No, m.NIM, m.Nama, m.IPK)
	}
}
