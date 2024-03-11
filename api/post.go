package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/m8-XD/real_estate_selling_website/db/repository"
	temp "github.com/m8-XD/real_estate_selling_website/templ"

	"github.com/m8-XD/real_estate_selling_website/db/models"
	"github.com/m8-XD/real_estate_selling_website/logging"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Header().Add("HX-Redirect", "/error?code=5xx")
		return
	}

	header := r.PostFormValue("header")
	desc := r.PostFormValue("desc")
	name := r.PostFormValue("name")
	phone := r.PostFormValue("phone")
	sPrice := r.PostFormValue("price")
	shType := r.PostFormValue("type")
	rooms := r.PostFormValue("rooms")
	mail := r.PostFormValue("email")
	pass := r.PostFormValue("pass")

	price, err := strconv.ParseInt(sPrice, 10, 64)
	if err != nil {
		w.Header().Add("HX-Redirect", "/error?code=5xx")
		return
	}

	hType := repository.HouseTypeId(shType)

	post := models.REstate{
		Header:      header,
		Description: desc,
		Name:        name,
		Phone:       phone,
		Price:       price,
		Rooms:       rooms,
		Type:        hType,
	}

	ok := post.Validate()
	if !ok {
		logging.Info("user entered wrong something")
		templ.Handler(temp.ValidatePostCreation(post, mail, false)).ServeHTTP(w, r)
		return
	}

	logging.Info("%v", post)

	ok, id := repository.GetId(mail, pass)
	if !ok {
		logging.Info("user entered wrong password")
		templ.Handler(temp.ValidatePostCreation(post, mail, false)).ServeHTTP(w, r)
		return
	}

	post.CreatorId = id

	postId, err := repository.CreatePost(&post)
	if err != nil {
		logging.Error("there was an error creating the post")
		logging.Error(err.Error())
		w.Header().Add("HX-Redirect", "/error?code=5xx")
		return
	}
	logging.Info("creating post %v", post)
	w.Header().Add("HX-Redirect", fmt.Sprintf("/post/%d", postId))
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	//map[1:[on] 2:[on] 5+:[on] Cottage:[on] Flat:[on] order:[recent] page:[1]]
	var err error

	q := r.URL.Query()
	var page int64

	if q.Has("page") {
		page, err = strconv.ParseInt(q.Get("page"), 10, 64)
	} else {
		page = 1
	}
	if err != nil {
		w.Header().Add("HX-Redirect", "/error?code=404")
	}

	posts := repository.AllPostsByFilter(map[string][]string{"type": {repository.HouseTypeId("Flat")}}, "updated_at", 1)

	hasNext := len(posts) == repository.PageSize
	hasPrev := page > 1

	templ.Handler(temp.AllPosts(posts, []string{"1", "2", "5+", "Other", "recent"}, hasNext, hasPrev)).ServeHTTP(w, r)
}
