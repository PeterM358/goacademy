package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Email    string `header:"Email" gorm:"unique"`
	Password string `header:"Password"`
	FirstName string `header:"First Name" gorm:"default:Change name"`
	LastName  string `header:"Last Name" gorm:"default:Change name"`
	Vehicles  []Vehicle
	Mechanics []Mechanic
}

type Mechanic struct {
	gorm.Model
	OwnerID uint
	CompanyName string    `header:"Company Name"`
	Address     string    `header:"Address"`
	Phone       uint      `header:"Phone"`
	Description string    `header:"Description"`
}

type Vehicle struct {
	gorm.Model
	OwnerID uint
	Brand      string `header:"Brand"`
	BrandModel string `header:"Brand Model"`
	Repairs    []Repair
}

type Repair struct {
	gorm.Model
	VehicleID uint
	Title       string `header:"Title"`
	Description string `header:"Description"`
	Price       int    `header:"Price"`
}

//TODO password auth
//func Hash(password string) ([]byte, error) {
//	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//}
//
//func VerifyPassword(hashedPassword, password string) error {
//	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
//}
