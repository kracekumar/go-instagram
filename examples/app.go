package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kracekumar/go-instagram"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	i := instagram.NewInstagram(CLIENT_ID, CLIENT_SECRET)
	resp := i.Authenticate(REDIRECT_URL, "") //PAss empty scope
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(rbody))
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.FormValue("code"))
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rbody)
	fmt.Fprintf(w, string(rbody))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/redirect/instagram", RedirectHandler)
	http.Handle("/", r)
	http.ListenAndServe(":9999", nil)
}
