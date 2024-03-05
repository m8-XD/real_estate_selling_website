package main

import (
	"net/http"

	"github.com/m8-XD/real_estate_selling_website/api/user_api"
	"github.com/m8-XD/real_estate_selling_website/db/repository"
	"github.com/m8-XD/real_estate_selling_website/logging"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/testdb", user_api.CreateUser)

	logging.Ok("server has started")
	_ = http.ListenAndServe(":8080", mux)
}
