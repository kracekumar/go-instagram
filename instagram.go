package instagram

import (
	//	"errors"
	"code.google.com/p/goauth2/oauth"
	//	"io/ioutil"
	"log"
	"net/http"

//	"net/url"
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
	AccessToken string
	client      *http.Client
	jar         *Jar
	Code        string
	OauthConfig *oauth.Config
}

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

/*func (i Instagram) resolve_redirect(resp *http.Response) *http.Response {
	for count := 0; count < REDIRECT_LIMIT; count++ {
		log.Println(resp.StatusCode)
		if resp.StatusCode == 200 || resp.StatusCode >= 400 {
			return resp
		} else {
			for _, code := range REDIRECT_CODES {
				if code == resp.StatusCode {
					location, err := resp.Location()
					check_error(err)
					//method := resp.Request.Method
					//				log.Println(resp.Header)
					//rbody, err := ioutil.ReadAll(resp.Body)
					//log.Println(resp.Request.Form)
					//check_error(err)
					//log.Println(string(rbody))
					//log.Println(location)
					resp, err := i.client.PostForm(location.String(), i.URLValues)
					log.Println(resp)
					log.Println(resp.StatusCode)
					break
				}
			}

		}
	}
	panic("MAximim Redirects")
	return resp
}*/

func NewInstagram(client_id, client_secret string) *Instagram {
	i := &Instagram{}
	i.OauthConfig = &oauth.Config{
		ClientId:     client_id,
		ClientSecret: client_secret,
	}
	i.client = &http.Client{}
	return i
}

func (i Instagram) Authenticate(redirect_url, scope string) string {
	//first step: Direct user to Instagram authorization URL
	i.OauthConfig.AuthURL = AUTHORIZATION_URL
	i.OauthConfig.TokenURL = ACCESS_URL
	i.OauthConfig.RedirectURL = redirect_url
	i.OauthConfig.Scope = scope
	return i.OauthConfig.AuthCodeURL("")
	//i.OauthConfig.AuthCodeURL("")
}

func (i Instagram) GetAccessToken(code string) {
	//Third step of oauth2.0: Do a Post
	t := &oauth.Transport{oauth.Config: i.OauthConfig}
	t.Exchange(code)
	log.Println(t)
	/*check_error(err)
	resp, err := i.client.PostForm(u.String(), url.Values{"client_id": {i.ClientId}, "client_secret": {i.ClientSecret},
		"redirect_uri": {i.RedirectURI}, "code": {code}, "grant_type": {"authorization_code"}})
	check_error(err)
	log.Println("Get Access Token")
	log.Println(resp)*/
}
