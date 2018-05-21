package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	discordAPI             = "https://discordapp.com/api"
	discordOauth2          = discordAPI + "/oauth2"
	discordOauth2Token     = discordOauth2 + "/token"
	discordOauth2Authorize = discordOauth2 + "/authorize"
)

var (
	discordTokenRedirect = discordOauth2Authorize + "?response_type=code&client_id=" + clientID + "&scope=identify&redirect_uri=" + url.QueryEscape(base+"/api/login/callback")
	discordAuthToken     string
	errInvalidToken      = errors.New("Invalid code")
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func getTokenAuthHeader() string {
	if discordAuthToken == "" {
		discordAuthToken = "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
	}
	return discordAuthToken
}

func getToken(code string) (*tokenResponse, error) {
	client := &http.Client{}

	println(discordOauth2Token + "?grant_type=authorization_code&code=" + code)
	req, err := http.NewRequest("POST", discordOauth2Token+"?grant_type=authorization_code&code="+code, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", getTokenAuthHeader())
	fmt.Printf("%+v\n", req)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	println(res.Status)
	println(string(body))
	if res.StatusCode != http.StatusOK {
		return nil, errInvalidToken
	}

	var tr *tokenResponse
	json.Unmarshal(body, &tr)

	return tr, nil
}
