package main

import "fmt"

func ArrayFunc() {

	// Array - fixed data (integer & string)
	Numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println(Numbers)

	Strings := [...]string{"PHP", "Golang", "NodeJS"}
	fmt.Println(Strings)

}

func SliceFunc() {

	// Slice - dynamic data
	Makanan := make([]string, 3)
	Makanan = append(Makanan, "Nasi")
	Makanan = append(Makanan, "Ikan", "Tempe")
	fmt.Println(Makanan)

}

func MapFunc() {

	// Map - key value
	VehiclePlatCode := map[string]string{
		"B":  "Jakarta",
		"D":  "Bandung",
		"F":  "Bogor",
		"DK": "Bali",
	}

	fmt.Println(VehiclePlatCode)

}

type VehicleData struct {
	Name           string
	Brand          string
	Stock          int
	StockAvailable bool
}

func StructFunc() {

	// Struct - with string, integer, boolean
	Vehicles := []VehicleData{
		{
			Name:           "HRV",
			Brand:          "Honda",
			Stock:          5,
			StockAvailable: true,
		},
		{
			Name:           "Innova Zenix",
			Brand:          "Toyota",
			Stock:          0,
			StockAvailable: false,
		},
	}

	fmt.Println(Vehicles)

}

func InterfaceFunc() {

	Vehicles := map[string]interface{}{
		"Brand":     "Toyota",
		"Name":      "yaris Cross",
		"Price":     450000000,
		"Available": true,
		"Stock":     50,
	}

	fmt.Println(Vehicles)

}

func main() {

	ArrayFunc()
	SliceFunc()
	StructFunc()
	InterfaceFunc()

}
