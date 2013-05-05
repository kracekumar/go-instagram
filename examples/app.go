package main

import (
	"github.com/gorilla/mux"
	"instagram"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	i := NewInstagram(CLIENT_ID, CLIENT_SECRET)
	i.Authenticate(REDIRECT_URL, "") //PAss empty scope
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/redirect/instagram", RedirectHandler)
	http.Handle("/", r)
	http.ListenAndServe(":9999", nil)
}
