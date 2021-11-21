package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `header:"Email"`
	Password string `header:"Password"`
	//OwnerRef uint   `header:"Owner Ref"`
	Owner    Owner
	Mechanic Mechanic
	//MechanicRef uint     `header:"Mechanic Ref"`
	//Mechanic    Mechanic `gorm:"foreignKey:Mechanic Ref"`
}

// should create boolean to add either owner or mechanic profile TODO
// both profiles can manage same vehicles struct if permitted TODO

type Owner struct {
	gorm.Model
	UserID uint
	//UserID    uint      `gorm:"foreignKey:OwnerRef"`
	FirstName string `header:"First Name" gorm:"default:Change name"`
	LastName  string `header:"Last Name" gorm:"default:Change name"`
	Vehicles  []Vehicle
}

type Mechanic struct {
	gorm.Model
	UserID      uint
	CompanyName string    `header:"Company Name"`
	Address     string    `header:"Address"`
	Phone       uint      `header:"Phone"`
	Description string    `header:"Description"`
	Vehicles    []Vehicle

}

type Vehicle struct {
	gorm.Model
	ID         uint
	OwnerID    uint
	MechanicID uint
	Brand      string `header:"Brand"`
	BrandModel string `header:"Brand Model"`
	Repairs    []Repair
}

type Repair struct {
	VehicleID   uint
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
