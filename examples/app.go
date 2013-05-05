package main

import (
	"github.com/gorilla/mux"
	"github.com/kracekumar/go-instagram"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	i := instagram.NewInstagram(CLIENT_ID, CLIENT_SECRET)
	i.Authenticate(REDIRECT_URL, "") //PAss empty scope
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Form)
	log.Println(r.FormValue("code"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/redirect/instagram", RedirectHandler)
	http.Handle("/", r)
	http.ListenAndServe(":9999", nil)
}
