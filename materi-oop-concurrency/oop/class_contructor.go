package oop

import "fmt"

type Mahasiswa struct {
	ID int
	Nama string
	NIM string
	Umur int
	Nilai float64
}

func NewStudent(id int, nama string, nim string, umur int) *Mahasiswa {
	// bagian ini
	return &Mahasiswa{
		ID: id,
		Nama: nama,
		Umur: umur,
		NIM: nim,
		Nilai: 0.0,
	}
}

func (s *Mahasiswa) Study(hours float64) {
	s.Nilai += hours * 0.5 // cara 1
	// s.Nilai = s.Nilai + hours * 0.5 // cara 2

	if s.Nilai > 100 {
		s.Nilai = 100
	}
}

func (s Mahasiswa) GetInfo() string {
	return fmt.Sprintf("Mahasiswa %s (%s) - Nilai: %.2f", s.Nama, s.NIM, s.Nilai)
}

func GoConstructor(id int, name string, nim string, umur int) {
	mahasiswa := NewStudent(id, name, nim, umur)
	mahasiswa2 := &Mahasiswa{ID: id, Nama: name, NIM: nim, Umur: umur}
	mahasiswa.Study(10)
	mahasiswa2.Study(20)

	fmt.Println(mahasiswa.GetInfo())
	fmt.Println(mahasiswa2.GetInfo())

}
