package oauth

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"

	"encoding/base64"

	"discord.land/site/server/database"
	"discord.land/site/server/utilities/http"
)

// TokenResponse represents the received data from AuthorizeCode
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

var meEndpoint = discordgo.EndpointUser("@me")

func codeEndpoint(code string) string {
	return discordgo.EndpointOauth2 + "token?grant_type=authorization_code&code=" + code
}

// OAuth represents the data belonging to the app
type OAuth struct {
	AuthorizationHeader string
	ClientID            string
	ClientSecret        string
}

// New creates a new instance of OAuth
func New(clientID, clientSecret string) *OAuth {
	return &OAuth{
		AuthorizationHeader: "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)),
		ClientID:            clientID,
		ClientSecret:        clientSecret,
	}
}

// AuthorizeCode authorizes a code
func (o *OAuth) AuthorizeCode(code string) (tr *TokenResponse, err error) {
	body, err := http.Request("POST", codeEndpoint(code), http.Headers{"Authorization": o.AuthorizationHeader})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &tr)
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// FetchUser gets a user's information by their access token
func (o *OAuth) FetchUser(accessToken string) (du *database.DiscordUser, err error) {
	body, err := http.Request("GET", meEndpoint, http.Headers{"Authorization": "Bearer " + accessToken})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &du)
	if err != nil {
		return nil, err
	}

	return du, nil
}
