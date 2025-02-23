package main

import (
	"fmt"
	"time"
)

// Struct untuk menyimpan biodata mahasiswa
type Mahasiswa struct {
	Nama    string
	Usia    int
	Alamat  string
}

// Interface untuk hewan dengan metode berbicara
type Hewan interface {
	Bersuara() string
}

// Struct Anjing yang mengimplementasikan interface Hewan
type Anjing struct{}

func (a Anjing) Bersuara() string {
	return "Guk! Guk!"
}

// Struct Kucing yang mengimplementasikan interface Hewan
type Kucing struct{}

func (k Kucing) Bersuara() string {
	return "Meong!"
}

// Node untuk Linked List
type Simpul struct {
	Nilai int
	Lanjut *Simpul
}

// Struktur LinkedList
type DaftarTertaut struct {
	Kepala *Simpul
}

// Menambahkan nilai baru ke linked list
func (dt *DaftarTertaut) Tambah(nilai int) {
	baru := &Simpul{Nilai: nilai}
	if dt.Kepala == nil {
		dt.Kepala = baru
		return
	}
	penunjuk := dt.Kepala
	for penunjuk.Lanjut != nil {
		penunjuk = penunjuk.Lanjut
	}
	penunjuk.Lanjut = baru
}

// Menampilkan isi linked list
func (dt *DaftarTertaut) Tampilkan() {
	penunjuk := dt.Kepala
	for penunjuk != nil {
		fmt.Print(penunjuk.Nilai, " -> ")
		penunjuk = penunjuk.Lanjut
	}
	fmt.Println("nil")
}

// Struktur Queue (Antrian)
type Antrian struct {
	elemen []int
}

// Menambahkan elemen ke dalam antrian
func (a *Antrian) Tambah(item int) {
	a.elemen = append(a.elemen, item)
}

// Menghapus elemen dari antrian
func (a *Antrian) Hapus() (int, bool) {
	if len(a.elemen) == 0 {
		return 0, false // Antrian kosong
	}
	item := a.elemen[0]
	a.elemen = a.elemen[1:] // Hapus elemen pertama
	return item, true
}

// Menampilkan isi antrian
func (a *Antrian) Tampilkan() {
	fmt.Println("Antrian:", a.elemen)
}

// Struktur Stack (Tumpukan)
type Tumpukan struct {
	elemen []int
}

// Menambahkan elemen ke dalam stack
func (t *Tumpukan) Tambah(item int) {
	t.elemen = append(t.elemen, item)
}

// Menghapus elemen dari stack
func (t *Tumpukan) Hapus() (int, bool) {
	if len(t.elemen) == 0 {
		return 0, false // Stack kosong
	}
	item := t.elemen[len(t.elemen)-1]
	t.elemen = t.elemen[:len(t.elemen)-1] // Hapus elemen terakhir
	return item, true
}

// Menampilkan isi stack
func (t *Tumpukan) Tampilkan() {
	fmt.Println("Tumpukan:", t.elemen)
}

func main() {
	// Contoh penggunaan array dan slice
	namaMahasiswa := [2]string{"Alandrian", "Surya"}
	nilaiMahasiswa := [4]int{80, 85, 90, 95}
	kehadiranMahasiswa := [4]bool{true, false, true, true}
	ipkMahasiswa := [4]float64{3.5, 3.7, 3.9, 4.0}

	// Slice dengan nilai tambahan
	skorTugas := []int{75, 85}
	skorTugas = append(skorTugas, 95) // Menambahkan nilai baru

	// Map daftar mahasiswa dengan NIM sebagai kunci
	daftarMahasiswa := map[string]int{
		"21001": 22,
		"21002": 21,
		"21003": 23,
	}

	// Menampilkan informasi
	fmt.Println("Nama Mahasiswa:", namaMahasiswa[0])
	fmt.Println("Nilai Mahasiswa:", nilaiMahasiswa[0])
	fmt.Println("Kehadiran Mahasiswa:", kehadiranMahasiswa[0])
	fmt.Println("IPK Mahasiswa:", ipkMahasiswa[0])
	fmt.Println("Skor Tugas:", skorTugas[2])
	fmt.Println("Usia Mahasiswa dengan NIM 21001:", daftarMahasiswa["21001"])

	// Membuat biodata mahasiswa
	mahasiswa1 := Mahasiswa{
		Nama:   "Alice",
		Usia:   21,
		Alamat: "Jl. Merdeka No.10",
	}

	fmt.Println("\nBiodata Mahasiswa:", mahasiswa1)
	fmt.Println("Nama:", mahasiswa1.Nama)
	fmt.Println("Usia:", mahasiswa1.Usia)
	fmt.Println("Alamat:", mahasiswa1.Alamat)

	// Mengubah usia mahasiswa
	mahasiswa1.Usia = 22
	fmt.Println("Usia setelah diperbarui:", mahasiswa1.Usia)

	// Contoh penggunaan pointer
	angka := 10
	ptr := &angka

	fmt.Println("\nNilai angka:", angka)
	fmt.Println("Alamat angka:", ptr)

	// Mengubah nilai melalui pointer
	*ptr = 20
	fmt.Println("Nilai setelah diubah melalui pointer:", angka)

	// Contoh penggunaan interface Hewan
	hewanPeliharaan := []Hewan{Anjing{}, Kucing{}}
	for _, hewan := range hewanPeliharaan {
		fmt.Println(hewan.Bersuara())
	}

	// Contoh penggunaan goroutine dan channel
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Second) // Simulasi delay 1 detik
		}
		close(ch) // Tutup channel setelah selesai
	}()

	// Menerima data dari channel
	fmt.Println("\nMenerima data dari channel:")
	for nilai := range ch {
		fmt.Println("Diterima:", nilai)
	}

	// Contoh penggunaan Linked List
	ll := &DaftarTertaut{}
	ll.Tambah(1)
	ll.Tambah(2)
	ll.Tambah(3)
	fmt.Println("\nLinked List:")
	ll.Tampilkan()

	// Contoh penggunaan Queue (Antrian)
	antrian := &Antrian{}
	antrian.Tambah(1001)
	antrian.Tambah(1002)
	antrian.Tambah(1003)
	fmt.Println("\nAntrian Mahasiswa:")
	antrian.Tampilkan()

	item, _ := antrian.Hapus()
	fmt.Println("Mahasiswa dengan NIM", item, "diproses.")
	antrian.Tampilkan()

	// Contoh penggunaan Stack (Tumpukan)
	tumpukan := &Tumpukan{}
	tumpukan.Tambah(2001)
	tumpukan.Tambah(2002)
	tumpukan.Tambah(2003)
	fmt.Println("\nTumpukan Mahasiswa:")
	tumpukan.Tampilkan()

	item, _ = tumpukan.Hapus()
	fmt.Println("Mahasiswa dengan NIM", item, "diambil dari tumpukan.")
	tumpukan.Tampilkan()
}
