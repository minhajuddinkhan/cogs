package cogs

import (
	"fmt"
	"strings"

	"github.com/minhajuddinkhan/cogs/ciphers"
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/minhajuddinkhan/cogs/types"
)

const (
	PrivateKey = "pv_key"
	PublicKey  = "pub_key"
	Bucket     = "cipher"
	CredsKey   = "credentials"
	Delimiter  = "|"
)

func Update(store bolt.Store, c *types.Credentials) error {

	var (
		accToken string
		hash     []byte
	)
	pvKey, err := ciphers.BytesToPrivateKey(c.PrivateKey)
	if err != nil {
		return err
	}

	text, err := ciphers.DecryptWithPrivateKey(c.Hash, pvKey)
	if err != nil {
		return err
	}

	creds := strings.Split(string(text), Delimiter)
	accToken, err = GetAccessToken(creds[0], creds[1])
	if err != nil {
		valid := false
		var username, pwd string
		for !valid {
			username, pwd, err := GetUserAndPwd()
			if err != nil {
				return err
			}

			accToken, err = GetAccessToken(username, pwd)
			if err != nil {
				if "y" == TakeInput("Invalid Credentials. Do you want to continue? [y/N]") {
					continue
				}
				return fmt.Errorf("Invalid credentials")
			}
			valid = true

		}
		text := username + Delimiter + pwd
		pubKey, err := ciphers.BytesToPublicKey(c.PublicKey)
		if err != nil {
			return err
		}
		hash, err = ciphers.EncryptWithPublicKey([]byte(text), pubKey)
		if err != nil {
			return err
		}
	}

	c.Hash = hash
	c.AccessToken = accToken
	return store.Create([]byte(CredsKey), c, Bucket)
}
