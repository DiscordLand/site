package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"discord.land/site/server/database"
)

const (
	discordAPI               = "https://discordapp.com/api"
	discordOauth2            = discordAPI + "/oauth2"
	discordOauth2Authorize   = discordOauth2 + "/authorize"
	discordOauth2Token       = discordOauth2 + "/token"
	discordOauth2TokenRevoke = discordOauth2Token + "/revoke"

	contentType = "application/x-www-form-urlencoded"
)

var (
	discordOauth2URL      = discordOauth2Authorize + "?client_id=" + clientID + "&callback_uri=" + url.QueryEscape(host+"/api/login/callback")
	discordOauth2Refresh  = discordOauth2URL + "&client_secret=" + clientSecret + "&grant_type=refresh_token&refresh_token="
	discordOauth2Redirect = discordOauth2URL + "&response_type=code&scope=identify"
	discordOauth2Header   = "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))

	errInvalidAccessToken = errors.New("Invalid access token")
	errInvalidOauth2Code  = errors.New("Invalid code")
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func authorizeOauth2Code(code string) (*tokenResponse, error) {
	req, err := http.NewRequest("POST", discordOauth2Token+"?grant_type=authorization_code&code="+code, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", discordOauth2Header)
	req.Header.Set("Content-Type", contentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errInvalidOauth2Code
	}

	var tr *tokenResponse
	err = json.Unmarshal(body, &tr)
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// func refreshOauth2Token(token string) (*tokenResponse, error) {
// 	res, err := http.Post(discordOauth2Refresh+token, contentType, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	println(string(body))
// 	println(res.Status)
// 	if res.StatusCode != http.StatusOK {
// 		return nil, errInvalidAccessToken
// 	}

// 	var tr *tokenResponse
// 	err = json.Unmarshal(body, &tr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tr, nil
// }

func fetchDiscordUser(accessToken string) (*database.DiscordUser, error) {
	req, err := http.NewRequest("GET", discordAPI+"/users/@me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errInvalidAccessToken
	}

	var du *database.DiscordUser
	err = json.Unmarshal(body, &du)
	if err != nil {
		return nil, err
	}

	return du, nil
}
