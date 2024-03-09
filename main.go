package main

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/m8-XD/real_estate_selling_website/api"
	"github.com/m8-XD/real_estate_selling_website/logging"
	temp "github.com/m8-XD/real_estate_selling_website/templ"
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
		code := r.URL.Query().Get("code")
		if !r.URL.Query().Has("code") {
			code = "404"
		}
		if strings.EqualFold(code, "5xx") {
			http.ServeFile(rw, r, "./static/error/5xx.html")
		} else {
			templ.Handler(temp.ShowError(code)).ServeHTTP(rw, r)
		}
	})
	mux.HandleFunc("/post/create", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./static/post/create.html")
	})
	mux.HandleFunc("/post/create/create.css", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./static/post/create.css")
	})
	mux.HandleFunc("/post/all", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./static/post/all.html")
	})

	mux.HandleFunc("POST /api/v1/user/email", api.ValidateEmail)
	// mux.HandleFunc("POST /api/v1/user/delete/{id}")
	// mux.HandleFunc("POST /api/v1/user/edit/{id}")
	mux.HandleFunc("POST /api/v1/user/create", api.CreateUser)
	mux.HandleFunc("POST /api/v1/post/create", api.CreatePost)
	mux.Handle("/api/v1/post/form", templ.Handler(temp.PostCreationForm()))
	// mux.HandleFunc("POST /api/v1/post/delete/{id}")
	// mux.HandleFunc("POST /api/v1/post/edit/{id}")
	// mux.HandleFunc("/post/{id}")

	logging.Ok("server has started")
	_ = http.ListenAndServe(":8080", mux)
}
