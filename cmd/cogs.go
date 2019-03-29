package cmd

import (
	"fmt"
	"log"
	"syscall"

	"github.com/minhajuddinkhan/cogs/ciphers"
	"github.com/minhajuddinkhan/cogs/types"

	"github.com/minhajuddinkhan/cogs"
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	PrivateKey = "pv_key"
	PublicKey  = "pub_key"
	Bucket     = "cipher"
	CredsKey   = "credentials"
)

func hasLoggedInBefore(store bolt.Store) bool {
	var creds types.Credentials
	err := store.Get([]byte(CredsKey), &creds, Bucket)
	return err == nil
}

const Delimiter = "|"

// BeforeAction BeforeActionu
var BeforeAction = func(store bolt.Store) cli.BeforeFunc {
	return func(c *cli.Context) error {
		if hasLoggedInBefore(store) {
			return nil
		}
		username := cogs.TakeInput("Enter your username")
		fmt.Println("Enter Password")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		password := string(bytePassword)

		_, err = cogs.GetAccessToken(username, password)
		if err != nil {
			return err
		}

		pvKey, pubKey, err := ciphers.GenerateKeyPair(1500)
		if err != nil {
			return err
		}

		text := username + Delimiter + password
		hash, err := ciphers.EncryptWithPublicKey([]byte(text), pubKey)
		if err != nil {
			return err
		}
		creds := types.Credentials{
			PublicKey: func() []byte {
				b, err := ciphers.PublicKeyToBytes(pubKey)
				if err != nil {
					log.Fatal(err)
				}
				return b
			}(),
			PrivateKey: ciphers.PrivateKeyToBytes(pvKey),
			Hash:       hash,
		}

		return store.Create([]byte(CredsKey), creds, Bucket)
	}

}
