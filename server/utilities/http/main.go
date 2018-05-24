package http

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Headers to set for the request
type Headers map[string]string

// Request a URL and return its body
func Request(method, url string, headers Headers) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// SendBadRequest sends an error 400 along with a JSON message
func SendBadRequest(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"message": message})
}
