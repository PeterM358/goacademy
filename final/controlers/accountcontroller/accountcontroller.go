package accountcontroller

import (
	"final/models"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"text/template"
)

func Index(response http.ResponseWriter, request *http.Request){
	tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
	tmp.Execute(response, nil)
}

//func Register(response http.ResponseWriter, request *http.Request){
//	tmp, _ := template.ParseFiles("views/accountcontroller/register.html")
//	tmp.Execute(response, nil)
//}

var tmpl = template.Must(template.ParseFiles("views/accountcontroller/register.html"))
func Register(response http.ResponseWriter, request *http.Request){
	//dsn := "root:root@tcp(127.0.0.1:3306)/dbFinal?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal(err)
	//}

	http.HandleFunc("/account/register/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		email := request.FormValue("Email")
		password := request.FormValue("Password")
		firstName := request.FormValue("First Name")
		lastName := request.FormValue("Last Name")
		user := &models.Owner{
			Model:     gorm.Model{},
			Email:     email,
			Password:  password,
			FirstName: firstName,
			LastName:  lastName,
			Vehicles:  nil,
			Mechanics: nil,
		}
		tmpl.Execute(response, struct {Success bool}{true})
		fmt.Println(user)

	})

	//nu := db.Create(user)
	//if nu.Error != nil {
	//	http.Error(response, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//http.Redirect(response, request, "/", http.StatusFound)
}



func Login(response http.ResponseWriter, request *http.Request){
	request.ParseForm()
	email := request.Form.Get("Email")
	password := request.Form.Get("Password")
	if email == "koki@abv.bg" && password == "asdf123"{
		data := map[string]interface{} {
			"err": "Invalid",
		}
		tmp, _ := template.ParseFiles("views/accountcontroller/login.html")
		tmp.Execute(response, data)
	} else {

		http.Redirect(response, request, "index", http.StatusSeeOther )
	}
}