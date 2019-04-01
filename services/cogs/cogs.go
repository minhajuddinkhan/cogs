package cogs

import (
	"fmt"
	"strings"

	"github.com/minhajuddinkhan/cogs/ciphers"
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/minhajuddinkhan/cogs/types"
)

const (
	// PrivateKey PrivateKey
	PrivateKey = "pv_key"
	PublicKey  = "pub_key"
	Bucket     = "cipher"
	CredsKey   = "credentials"
	Delimiter  = "|"
)

// Cogs cogs interface
type Cogs interface {
	Update() error
	Credentials() *types.Credentials
}
type cogsHandler struct {
	Creds *types.Credentials
	Store bolt.Store
}

func (c *cogsHandler) Credentials() *types.Credentials {
	return c.Creds
}

// New new cogs handler
func New(creds *types.Credentials, store bolt.Store) Cogs {
	return &cogsHandler{Creds: creds, Store: store}
}

func (c *cogsHandler) Update() error {

	pvKey, err := ciphers.BytesToPrivateKey(c.Creds.PrivateKey)
	if err != nil {
		return fmt.Errorf("error converting bytes to private key. err: %v", err)
	}
	text, err := ciphers.DecryptWithPrivateKey(c.Creds.Hash, pvKey)
	if err != nil {
		return fmt.Errorf("error decrypting with private key. err: %v", err)
	}

	creds := strings.Split(string(text), Delimiter)
	attrs, err := GetAccessToken(creds[0], creds[1])
	if err != nil {
		return err
	}
	c.Creds.AccessToken = attrs.AccessToken
	if err != nil {
		valid := false
		var username, pwd string
		for !valid {
			username, pwd, err := GetUserAndPwd()
			if err != nil {
				return err
			}

			attrs, err = GetAccessToken(username, pwd)
			c.Creds.AccessToken = attrs.AccessToken
			c.Creds.EmployeeID = attrs.EmployeeID
			if err != nil {
				if "y" == TakeInput("Invalid Credentials. Do you want to continue? [y/N]") {
					continue
				}
				return fmt.Errorf("Invalid credentials")
			}
			valid = true

		}
		text := username + Delimiter + pwd
		pubKey, err := ciphers.BytesToPublicKey(c.Creds.PublicKey)
		if err != nil {
			return err
		}
		c.Creds.Hash, err = ciphers.EncryptWithPublicKey([]byte(text), pubKey)
		if err != nil {
			return err
		}
	}

	return c.Store.Create([]byte(CredsKey), c.Creds, Bucket)
}
