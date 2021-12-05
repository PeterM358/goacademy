package main

import (
	"final/controlers/accountcontroller"
	"final/models"
	"flag"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var port = flag.String("port", "8080", "final service address")

func main() {
	//smth broke with docker TODO
	//dsn := "root:test@tcp(localhost:3306)/dbTest?charset=utf8&parseTime=True&loc=Local"

	dsn := "root:root@tcp(localhost:3306)/dbFinal?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Owner{})
	db.AutoMigrate(&models.Mechanic{})
	db.AutoMigrate(&models.Vehicle{})
	db.AutoMigrate(&models.Repair{})
	fmt.Println(1)


	ih := &accountcontroller.IndexHandler{}
	osuh := &accountcontroller.OwnerSignUpHandler{db}
	msuh := &accountcontroller.MechanicSignUpHandler{db}

	lh := &accountcontroller.LoginHandler{}

	mux := http.NewServeMux()

	mux.Handle("/", ih)
	mux.Handle("/owner/signup/", osuh)
	mux.Handle("/mechanic/signup/", msuh)
	mux.Handle("/account/signin/", lh)
	log.Fatal(http.ListenAndServe(":8080", mux))
	//
	//fmt.Printf("New user created with ID: %d -> %+v\nRows afffected: %d\n",
	//	user.ID, // returns inserted data's primary key
	//	user,
	//	result.RowsAffected, // returns inserted records count
	//)

}
