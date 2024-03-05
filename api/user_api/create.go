package user_api

import (
	"fmt"
	"net/http"

	"github.com/m8-XD/real_estate_selling_website/db/models"
	"github.com/m8-XD/real_estate_selling_website/db/repository"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	name := "test"
	mail := "test"
	id, err := repository.CreateUser(&models.User{Name: name, Mail: mail, Pass: "123"})
	if err != nil {
		_, _ = rw.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}

	_, _ = rw.Write([]byte(fmt.Sprintf("new id is %d", id)))
}
