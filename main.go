package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/m8-XD/real_estate_selling_website/api"
	"github.com/m8-XD/real_estate_selling_website/logging"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("cannot load .env file")
	}

	mux := http.NewServeMux()

	// mux.HandleFunc("/") //show all posts
	mux.HandleFunc("/register", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./static/auth/register.html")
	})
	mux.HandleFunc("/error", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./static/error/"+r.URL.Query().Get("code")+".html")
	})
	mux.HandleFunc("/post/create", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./static/post/create.html")
	})

	mux.HandleFunc("POST /api/v1/user/email", api.ValidateEmail)
	// mux.HandleFunc("POST /api/v1/user/delete/{id}")
	// mux.HandleFunc("POST /api/v1/user/edit/{id}")
	mux.HandleFunc("POST /api/v1/user/create", api.CreateUser)
	mux.HandleFunc("POST /api/v1/post/create", api.CreatePost)
	// mux.HandleFunc("POST /api/v1/post/delete/{id}")
	// mux.HandleFunc("POST /api/v1/post/edit/{id}")
	// mux.HandleFunc("/post/{id}")

	logging.Ok("server has started")
	_ = http.ListenAndServe(":8080", mux)
}
