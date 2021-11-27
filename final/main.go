package main

import (
	"final/controlers/accountcontroller"
	"final/models"
	"flag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "final service address")

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/dbFinal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Owner{})
	db.AutoMigrate(&models.Mechanic{})
	db.AutoMigrate(&models.Vehicle{})
	db.AutoMigrate(&models.Repair{})

	//user := models.Owner{
	//	Model:     gorm.Model{},
	//	Email:     "koki@abv.bg",
	//	Password:  "asdf",
	//	FirstName: "Koki",
	//	LastName:  "Kokov",
	//	Vehicles:  nil,
	//}
	//
	//result := db.Create(&user)
	//
	//if result.Error != nil {
	//	log.Fatal(result.Error)
	//}
	//
	//vehicle := models.Vehicle{
	//	Model:      gorm.Model{},
	//	OwnerID:    1,
	//	Brand:      "Ford",
	//	BrandModel: "Ka",
	//	Repairs:    nil,
	//}
	//
	//nv := db.Create(&vehicle)
	//
	//if nv.Error != nil {
	//	log.Fatal(nv.Error)
	//}
	http.HandleFunc("/", accountcontroller.Index)
	http.HandleFunc("/account/register/", accountcontroller.Register)
	http.HandleFunc("/account/login/", accountcontroller.Login)
	http.ListenAndServe(*addr, nil)
	//

	//fmt.Printf("New user created with ID: %d -> %+v\nRows afffected: %d\n",
	//	user.ID, // returns inserted data's primary key
	//	user,
	//	result.RowsAffected, // returns inserted records count
	//)




}
