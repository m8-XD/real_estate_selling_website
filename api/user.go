package api

import (
	"net/http"
	m "net/mail"

	"github.com/a-h/templ"
	"github.com/m8-XD/real_estate_selling_website/db/models"
	"github.com/m8-XD/real_estate_selling_website/db/repository"
	"github.com/m8-XD/real_estate_selling_website/logging"
	t "github.com/m8-XD/real_estate_selling_website/templ"
)

func ValidateEmail(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		templ.Handler(t.CheckMail("", "")).ServeHTTP(rw, r)
	}
	mail := r.PostFormValue("email")
	name := r.PostFormValue("name")
	templ.Handler(t.CheckMail(mail, name)).ServeHTTP(rw, r)
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		logging.Error("couldn't parse form to create user")
		rw.Header().Add("HX-Redirect", "/error?code=5xx")
	}

	mail := r.PostFormValue("email")
	name := r.PostFormValue("name")
	pass := r.PostFormValue("password")

	ok := repository.IsValidMail(mail)
	if !ok {
		logging.Info("someone trying to register existing email")
		templ.Handler(t.CheckMail(mail, name)).ServeHTTP(rw, r)
		return
	}

	_, err = m.ParseAddress(mail)
	if err != nil {
		logging.Error("user entered invalid mail")
		templ.Handler(t.CheckMail(mail, name)).ServeHTTP(rw, r)
		return
	}

	user := models.User{Name: name, Mail: mail, Pass: pass}
	err = repository.CreateUser(&user)
	if err != nil {
		logging.Error("couldn't create user")
		rw.Header().Add("HX-Redirect", "/error?code=5xx")
	}
	logging.Ok("new user has been created")
	rw.Header().Add("HX-Redirect", "/")
}
