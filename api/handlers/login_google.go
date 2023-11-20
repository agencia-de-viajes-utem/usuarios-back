package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var ssgolang *oauth2.Config
var RandomString = "randoms-tring"

func init() {

	err := godotenv.Load(".env.credentials")
	if err != nil {
		log.Fatal("Error loading .env.credentials file")
	}
	ssgolang = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
func LoginGoogle(w http.ResponseWriter, r *http.Request) {
	url := ssgolang.AuthCodeURL(RandomString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}
