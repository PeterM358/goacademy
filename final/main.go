package main

import (
	"final/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/dbFinal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Owner{})
	db.AutoMigrate(&models.Mechanic{})
	db.AutoMigrate(&models.Vehicle{})
	db.AutoMigrate(&models.Repair{})

	user := models.User{
		Model:    gorm.Model{},
		Email:    "user@abv.bg",
		Password: "asdf",
		Owner:    models.Owner{},
		Mechanic: models.Mechanic{},
	}

	result := db.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Printf("New user created with ID: %d -> %+v\nRows afffected: %d\n",
		user.ID, // returns inserted data's primary key
		user,
		result.RowsAffected, // returns inserted records count
	)

}
