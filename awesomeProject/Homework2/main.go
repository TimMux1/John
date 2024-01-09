package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	dsn := "user=postgres dbname=MyProject sslmode=disable password=1212"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migrate
	db.AutoMigrate(&User{})

	// Create
	user := User{Name: "John Doe", Email: "john@example.com"}
	db.Create(&user)
	fmt.Println("User created:", user)

	// Read
	var fetchedUser User
	db.First(&fetchedUser, user.ID)
	fmt.Println("User fetched by ID:", fetchedUser)

	// Update
	db.Model(&fetchedUser).Update("Name", "Doe John")
	fmt.Println("User updated:", fetchedUser)

	// Delete
	db.Unscoped().Delete(&fetchedUser, fetchedUser.ID)
	fmt.Println("User deleted:", fetchedUser)

	// Get all users
	var allUsers []User
	db.Find(&allUsers)
	fmt.Println("All users:", allUsers)

}
