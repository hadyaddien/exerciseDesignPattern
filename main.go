package main

import (
	"fmt"
)

// User interface dengan method GetInfo
type User interface {
	GetInfo() string
}

// Dosen struct untuk representasi data dosen
type Dosen struct {
	Nama   string
	NIP    string
	Matkul string
}

// Mahasiswa struct untuk representasi data mahasiswa
type Mahasiswa struct {
	Nama    string
	NIM     string
	Jurusan string
}

// Implementasi method GetInfo untuk Dosen
func (d Dosen) GetInfo() string {
	return fmt.Sprintf("Dosen - Nama: %s, NIP: %s, Mengajar: %s", d.Nama, d.NIP, d.Matkul)
}

// Implementasi method GetInfo untuk Mahasiswa
func (m Mahasiswa) GetInfo() string {
	return fmt.Sprintf("Mahasiswa - Nama: %s, NIM: %s, Jurusan: %s", m.Nama, m.NIM, m.Jurusan)
}

func main() {
	// Membuat objek Dosen
	dosen := Dosen{
		Nama:   "Khairul Abdi",
		NIP:    "12345",
		Matkul: "Pemrograman Backend",
	}

	// Membuat objek Mahasiswa
	mahasiswa := Mahasiswa{
		Nama:    "Hadyaddien",
		NIM:     "67890",
		Jurusan: "Sistem Informasi",
	}

	// Menggunakan method GetInfo untuk objek Dosen dan Mahasiswa
	fmt.Println(dosen.GetInfo())
	fmt.Println(mahasiswa.GetInfo())
}
