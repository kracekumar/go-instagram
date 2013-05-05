package instagram

import (
	//	"errors"
	//	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//API docs: http://instagram.com/developer

const (
	AUTHORIZATION_URL = "https://api.instagram.com/oauth/authorize"
	ACCESS_URL        = "https://api.instagram.com/oauth/access_token"
)

type Jar struct {
	cookies []*http.Cookie
}

type Instagram struct {
	ClientSecret string
	ClientId     string
	AccessToken  string
	RedirectURI  string
	client       *http.Client
	jar          *Jar
}

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func NewInstagram(client_id, client_secret string) *Instagram {
	i := &Instagram{}
	i.ClientId = client_id
	i.ClientSecret = client_secret
	i.client = &http.Client{}
	return i
}

func (i Instagram) Authenticate(redirect_uri, scope string) *http.Response {
	u, err := url.Parse(redirect_uri)
	check_error(err)
	i.RedirectURI = redirect_uri
	//first step: Direct user to Instagram authorization URL
	instagram_url, err := url.Parse(AUTHORIZATION_URL)
	check_error(err)
	var resp *http.Response
	if scope != "" {
		resp, err = i.client.PostForm(u.String(), url.Values{"client_id": {i.ClientId}, "redirect_uri": {i.RedirectURI},
			"response_type": {"code"}, "scope": {scope}})
	} else {
		resp, err = i.client.PostForm(u.String(), url.Values{"client_id": {i.ClientId}, "redirect_uri": {i.RedirectURI},
			"response_type": {"code"}})
	}
	check_error(err)
	return resp
}

func (i Instagram) GetAccessToken(code string) *http.Response {
	//Third step of oauth2.0: Do a Post
	u, err := url.Parse(ACCESS_URL)
	check_error(err)
	resp, err := i.client.PostForm(u.String(), url.Values{"client_id": {i.ClientId}, "client_secret": {i.ClientSecret},
		"redirect_uri": {i.RedirectURI}, "code": {code}, "grant_type": {"authorization_code"}})
	check_error(err)
	log.Println("Get Access Token")
	log.Println(resp)
	return resp
}
