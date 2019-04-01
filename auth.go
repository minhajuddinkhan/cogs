package cogs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type Attrs struct {
	AccessToken string `json:"access-token"`
	EmployeeID  int    `json:"employee-id"`
}
type AuthBody struct {
	Data struct {
		Attributes Attrs `json:"attributes"`
	} `json:"data"`
}

// GetAccessToken request for cogs
func GetAccessToken(username, password string) (*Attrs, error) {
	payload := `{"data":{"type":"auths","attributes":{"userName":"%s","password":"%s","keepMeLoggedIn":true}},"included":[]}`
	raw, err := Request(
		fmt.Sprintf("%s/auth/login", baseURL),
		http.MethodPost,
		fmt.Sprintf(payload, username, password),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var body AuthBody
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	spew.Dump(body.Data.Attributes.EmployeeID)
	return &body.Data.Attributes, nil
}
