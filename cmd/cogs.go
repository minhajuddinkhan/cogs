package cmd

import (
	"fmt"
	"log"
	"syscall"

	"github.com/minhajuddinkhan/cogs"
	"github.com/minhajuddinkhan/cogs/ciphers"
	"github.com/minhajuddinkhan/cogs/types"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/urfave/cli"
)

func hasLoggedInBefore(store bolt.Store, creds *types.Credentials) bool {
	err := store.Get([]byte(cogs.CredsKey), &creds, cogs.Bucket)
	return err == nil
}

// BeforeAction BeforeActionu
var BeforeAction = func(store bolt.Store, creds *types.Credentials) cli.BeforeFunc {
	return func(c *cli.Context) error {
		if (c.Args().Get(0)) == "" {
			return nil
		}
		if hasLoggedInBefore(store, creds) {
			return nil
		}
		validUsernameAndPassword := false
		var username, password, accessToken string
		var employeeID int
		for !validUsernameAndPassword {
			username = cogs.TakeInput("Enter your username")
			fmt.Println("Enter Password")
			bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			password = string(bytePassword)
			attrs, err := cogs.GetAccessToken(username, password)
			if err != nil {
				again := cogs.TakeInput("Invalid Credentials. Do you want to try again [y/N]")
				if again == "y" {
					continue
				}
				return cli.NewExitError("", 1)
			}
			accessToken = attrs.AccessToken
			employeeID = attrs.EmployeeID

			validUsernameAndPassword = true
		}

		pvKey, pubKey, err := ciphers.GenerateKeyPair(1500)
		if err != nil {
			return err
		}

		text := username + cogs.Delimiter + password
		hash, err := ciphers.EncryptWithPublicKey([]byte(text), pubKey)
		if err != nil {
			return err
		}
		creds.PublicKey = func() []byte {
			b, err := ciphers.PublicKeyToBytes(pubKey)
			if err != nil {
				log.Fatal(err)
			}
			return b
		}()

		creds.PrivateKey = ciphers.PrivateKeyToBytes(pvKey)
		creds.Hash = hash
		creds.AccessToken = accessToken
		creds.EmployeeID = employeeID
		return store.Create([]byte(cogs.CredsKey), creds, cogs.Bucket)
	}

}
