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
	url := i.Authenticate(REDIRECT_URL, "")
	log.Println(url)
	http.Redirect(w, r, url, 302)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("inside redirect")
	rbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(rbody))
	fmt.Fprintf(w, "%s", string(rbody))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/redirect/instagram", RedirectHandler)
	http.Handle("/", r)
	http.ListenAndServe(":9999", nil)
}
