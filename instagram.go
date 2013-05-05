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
	REDIRECT_LIMIT    = 30
)

var REDIRECT_CODES = []int{301, 302, 303, 307}

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

func (i Instagram) resolve_redirect(resp *http.Response) *http.Response {
	for count := 0; count < REDIRECT_LIMIT; count++ {
		if resp.StatusCode == 200 || resp.StatusCode >= 400 {
			return resp
		} else {
			for code := range REDIRECT_CODES {
				if code == resp.StatusCode {
					location, err := resp.Location()
					check_error(err)
					method := resp.Request.Method
					log.Println(resp)
					log.Println(method)
					log.Println(location)
				}
			}

		}
	}
	panic("MAximim Redirects")
	return resp
}

func NewInstagram(client_id, client_secret string) *Instagram {
	i := &Instagram{}
	i.ClientId = client_id
	i.ClientSecret = client_secret
	i.client = &http.Client{}
	return i
}

func (i Instagram) Authenticate(redirect_uri, scope string) {
	i.RedirectURI = redirect_uri
	//first step: Direct user to Instagram authorization URL
	instagram_url, err := url.Parse(AUTHORIZATION_URL)
	check_error(err)
	var resp *http.Response
	if scope != "" {
		resp, err = i.client.PostForm(instagram_url.String(), url.Values{"client_id": {i.ClientId}, "redirect_uri": {i.RedirectURI},
			"response_type": {"code"}, "scope": {scope}})
	} else {
		resp, err = i.client.PostForm(instagram_url.String(), url.Values{"client_id": {i.ClientId}, "redirect_uri": {i.RedirectURI},
			"response_type": {"code"}})
	}
	check_error(err)
	log.Println(resp)
	i.resolve_redirect(resp)
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
