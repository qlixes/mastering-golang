package main

import (
	"fmt"
	"mysql-go/config"
	"mysql-go/models"
)

func main() {
	// buat koneksi gorm dengan database
	db := config.ConnectDB()

	// buat migrations
	db.AutoMigrate(&models.User{}, &models.Role{})

	role := []models.Role{
		{Name: "Admin"},
		{Name: "Users"},
		{Name: "Manager"},
	}

	if err := db.CreateInBatches(role, 100).Error; err != nil {
		fmt.Println("Error membuat Role:", err)
		return
	}


	users := []models.User{
		{
			Name: "Rafi Admin",
			Email: "rafiAdmin@gmail.com",
			Password: "rafi123@",
			Roles: []models.Role{role[0]},
		},
		{
			Name: "Rafi Users",
			Email: "RafiManager@gmail.com",
			Password: "rafi123@",
			Roles: []models.Role{role[1]},
		},
		{
			Name: "Rafi Manager",
			Email: "RafiUsers@gmail.com",
			Password: "rafi123@",
			Roles: []models.Role{role[2]},
		},
		{
			Name: "Rafi All Role",
			Email: "RafiAll@gmail.com",
			Password: "rafi123@",
			Roles: []models.Role{role[0], role[1], role[2]},
		},
	}

	if err := db.CreateInBatches(users, 100).Error; err != nil {
		fmt.Println("Error membuat Users:", err)
		return
	}


	// if err := db.Create(&role).Error; err != nil {
	// 	fmt.Println("Error membuat role:", err)
	// 	return
	// }

	// user := models.User {
	// 	Name: "Taufik 2",
	// 	Email: "Taufik@gmail.com",
	// 	Password: "Taufik123@",
	// 	Roles: []models.Role{role},
	// }

	// if err := db.Create(&user).Error; err != nil {
	// 	fmt.Println("Error membuat role:", err)
	// 	return
	// }

	// create Users

	// // table 1
	// user := models.User{
	// 	Name: "Taufik",
	// 	Email: "taufikmulyawan@gmail.com",
	// 	Password: "taufik1321@",
	// }

	// // table 2
	// role := models.Role{
	// 	Name: "Admin",
	// }

	// result := db.Create(&user)

	// resultRole := db.Create(&role)

	// if result.Error != nil || resultRole.Error != nil {
	// 	fmt.Println("Error membuat user")
	// 	return
	// }

	fmt.Println("Membuat user berhasil")
}