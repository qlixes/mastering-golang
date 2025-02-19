package main

import (
	"fmt"
)

// menampilkan data dari sebuah array
func namaBuah(arr []string) {
	fmt.Println(arr)
}

// membuat slice dari array, dengan ketentuan ambil data indeks [0] sampai indeks [2]
func buahBesar(arr []string) {
	slicing := arr[:3]
	fmt.Println(slicing)
}

// menampilkan data dari sebuah Map
func buahMap(arr map[string]string) {
	fmt.Println(arr)
}

// menampilkan struct dari variable dataHafid
func dataHafid() {
	hafid := DataDiri{
		nama:             "Hafid Iqbal",
		alamat:           "Pekalongan",
		role:             "Go Developer",
		umur:             28,
		statusPernikahan: true,
	}
	fmt.Println(hafid)
}

// membuat struct DataDiri, yang berisikan nama, alamat, role, umur dan status pernikahan
type DataDiri struct {
	nama, alamat, role string
	umur               int
	statusPernikahan   bool
}

func main() {
	arrayBuah := [...]string{"Nanas", "Semangka", "Melon", "Pir", "Delima", "Jeruk", "Cherry"}

	namaBuah(arrayBuah[:]) // output : {"Nanas", "Semangka", "Melon", "Pir", "Delima", "Jeruk", "Cherry"}

	buahBesar(arrayBuah[:]) // output : [Nanas Semangka Melon]

	iconBuah := map[string]string{
		"Nanas":    "🍍",
		"Semangka": "🍉",
		"Melon":    "🍈",
		"Jeruk":    "🍊",
		"Cherry":   "🍒",
		"Pir":      "🍐",
	} // output : map[Cherry:🍒 Jeruk:🍊 Melon:🍈 Nanas:🍍 Pir:🍐 Semangka:🍉]

	buahMap(iconBuah)

}
