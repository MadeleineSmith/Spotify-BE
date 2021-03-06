package handlers

import (
	"net/http"
	"net/url"
	"os"
)

type LoginUserHandler struct {
	HTTPClient *http.Client
}

func (h LoginUserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	authorizeUrl, _ := url.Parse("https://accounts.spotify.com/authorize")
	scopes := "user-read-private user-read-email playlist-modify-private"

	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", os.Getenv("CLIENT_ID"))
	params.Add("scope", scopes)
	params.Add("redirect_uri", os.Getenv("REDIRECT_URI"))
	// TODO - add `state` query param

	authorizeUrl.RawQuery = params.Encode()

	http.Redirect(w, req, authorizeUrl.String(), http.StatusFound)
}
