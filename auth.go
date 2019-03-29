package cogs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthBody struct {
	Data AuthData `json:"data"`
}
type AuthData struct {
	Attributes AuthAttrs `json:"attributes"`
}
type AuthAttrs struct {
	AccessToken string `json:"access-token"`
}

// GetAccessToken request for cogs
func GetAccessToken(username, password string) (string, error) {
	payload := `{"data":{"type":"auths","attributes":{"userName":"%s","password":"%s","keepMeLoggedIn":true}},"included":[]}`
	raw, err := request(
		fmt.Sprintf("%s/auth/login", baseURL),
		http.MethodPost,
		fmt.Sprintf(payload, username, password),
		nil,
	)
	if err != nil {
		return "", err
	}

	var body AuthBody
	if err := json.Unmarshal(raw, &body); err != nil {
		return "", err
	}
	return body.Data.Attributes.AccessToken, nil
}
