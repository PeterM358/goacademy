package accountcontroller

import (
	"final/models"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"gorm.io/gorm"
)

var indexTemplate, _ = template.ParseFiles("views/accountcontroller/index.html")
var ownerSignUpTemplate, _ = template.ParseFiles("views/accountcontroller/ownersignup.html")
var MechanicSignUpTemplate, _ = template.ParseFiles("views/accountcontroller/mechanicsignup.html")

var singInTemplate, _ = template.ParseFiles("views/accountcontroller/signin.html")


type IndexHandler struct{}

func (*IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, nil)
}

type OwnerSignUpHandler struct {
	DB *gorm.DB
}

func (oh *OwnerSignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		ownerSignUpTemplate.Execute(w, nil)
	case "POST":
		fmt.Println(r)
		r.ParseMultipartForm(0)
		fmt.Printf("%+v", r.Form)
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		firstName := r.FormValue("First Name")
		lastName := r.FormValue("Last Name")
		user := &models.Owner{
			Model:     gorm.Model{},
			Email:     email,
			Password:  password,
			FirstName: firstName,
			LastName:  lastName,
			Vehicles:  nil,
			Mechanics: nil,
		}
		fmt.Printf("%+v", user)
		userReq := oh.DB.Create(user)
		if userReq.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusCreated)
		fmt.Printf("Created user: %+v", user)
	default:
		w.Write([]byte("Unknown method type"))
	}
}

//nu := db.Create(user)
//if nu.Error != nil {
//	http.Error(response, err.Error(), http.StatusInternalServerError)
//	return
//}
//http.Redirect(response, request, "/", http.StatusFound)

type MechanicSignUpHandler struct {
	DB *gorm.DB
}

func (mh *MechanicSignUpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		MechanicSignUpTemplate.Execute(w, nil)
	case "POST":
		fmt.Println(r)
		r.ParseMultipartForm(0)
		fmt.Printf("%+v", r.Form)
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		companyName := r.FormValue("Company Name")
		address := r.FormValue("Address")
		phone := r.FormValue("Phone")
		pn, _ := strconv.Atoi(phone)

		description := r.FormValue("Description")
		user := &models.Mechanic{
			Model:     gorm.Model{},
			Email:     email,
			Password:  password,
			CompanyName: companyName,
			Address:  address,
			Phone:  pn,
			Description:  description,
		}

	//	Model:     gorm.Model{},
	//	Email:     email,
	//	Password:  password,
	//	CompanyName: companyName,
	//	Address:  address,
	//	Phone:  pn,
	//	Description:  description,

		fmt.Printf("%+v", user)
		userReq := mh.DB.Create(user)
		if userReq.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusCreated)
		fmt.Printf("Created user: %+v", user)
	default:
		w.Write([]byte("Unknown method type"))
	}
}





type LoginHandler struct{}

func (*LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("Email")
	password := r.Form.Get("Password")
	if email == "koki@abv.bg" && password == "asdf123" {
		data := map[string]interface{}{
			"err": "Invalid",
		}

		singInTemplate.Execute(w, data)
	} else {

		http.Redirect(w, r, "index", http.StatusSeeOther)
	}
}

