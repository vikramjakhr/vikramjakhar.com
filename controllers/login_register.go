package controllers

import (
	"net/http"
	"github.com/choudhary92/vikramjakhar.com/models"
	"github.com/choudhary92/vikramjakhar.com/data"
	"strings"
	"fmt"
)

func init() {
	http.HandleFunc("/user/registeration", register)
	http.HandleFunc("/user/register", registration)
	http.HandleFunc("/user/join", join)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		Root(w, r)
		return
	}
	regInfo := models.RegistrationInfo{}
	regInfo.Email = models.Email(r.FormValue("email"))
	regInfo.FirstName = r.FormValue("firstName")
	regInfo.LastName = r.FormValue("lastName")
	regInfo.UserName = models.UserName(r.FormValue("userName"))
	regInfo.Password = models.Password(r.FormValue("password"))
	regInfo.ConfirmPassword = models.Password(r.FormValue("confPass"))
	//regInfo.Register()
	Root(w, r)
}

type JoinInfo struct {
	Email         string
	IsAlreadyUsed bool
}

func join(w http.ResponseWriter, r *http.Request) {
	var email string = ""
	if r.Method == "POST" {
		email = strings.TrimSpace(r.FormValue("email"))
	}
	fmt.Println(email)
	//templateInfo := data.TemplateInfo{Title:"Register", Name:"register-tmpl", Path:"views/user/register-tmpl.html"}
	//renderLayoutWithoutSidebar(w, r, templateInfo, templateInfo.Title, JoinInfo{Email:email, IsAlreadyUsed:false})
}

func registration(w http.ResponseWriter, r *http.Request) {
	var email string = ""
	if r.Method == "POST" {
		email = strings.TrimSpace(r.FormValue("email"))
	}
	fmt.Println(email)
	templateInfo := data.TemplateInfo{Title:"Register", Name:"register-tmpl", Path:"views/user/register-tmpl.html"}
	renderLayoutWithoutSidebar(w, r, templateInfo, templateInfo.Title, JoinInfo{Email:email, IsAlreadyUsed:false})
}

func login(w http.ResponseWriter, r *http.Request) {

}
