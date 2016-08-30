package controllers

import (
	"net/http"
	"strings"
	"github.com/choudhary92/vikramjakhar.com/data"
	"github.com/choudhary92/vikramjakhar.com/models"
)

func init() {
	http.HandleFunc("/user/register", register)
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
	Email string
	IsAlreadyUsed bool
}

func join(w http.ResponseWriter, r *http.Request) {
	var email string = ""
	if r.Method == "POST" {
		email = strings.TrimSpace(r.FormValue("email"))
	}
	templateInfo := data.TemplateInfoMap["register"]
	renderLayoutWithoutSidebar(w, r, templateInfo, templateInfo.Title, JoinInfo{Email:email, IsAlreadyUsed:false})
}

func login(w http.ResponseWriter, r *http.Request) {

}
