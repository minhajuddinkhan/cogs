package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/minhajuddinkhan/cogs"
	"github.com/minhajuddinkhan/cogs/ciphers"
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/minhajuddinkhan/cogs/types"
	"github.com/urfave/cli"
)

type lunchBody struct {
	Data []struct {
		Type       string `json:"type"`
		Attributes struct {
			MenuItem string `json:"menu-item"`
		} `json:"attributes"`
	} `json:"data"`
}

// Lunch Gets Lunch
func Lunch(store bolt.Store) cli.Command {
	return cli.Command{
		Name: "lunch",
		Action: func(c *cli.Context) error {
			var creds types.Credentials
			if err := store.Get([]byte(CredsKey), &creds, Bucket); err != nil {
				return err
			}
			pvKey, err := ciphers.BytesToPrivateKey(creds.PrivateKey)
			if err != nil {
				return err
			}
			text, err := ciphers.DecryptWithPrivateKey([]byte(creds.Hash), pvKey)
			if err != nil {
				return err
			}
			textCreds := strings.Split(string(text), Delimiter)
			token, err := cogs.GetAccessToken(textCreds[0], textCreds[1])
			if err != nil {
				return err
			}
			raw, err := cogs.Lunch(token)
			if err != nil {
				return err
			}
			var lb lunchBody
			err = json.Unmarshal(raw, &lb)
			if err != nil {
				return err
			}
			fmt.Printf("Todays lunch is %s \n", lb.Data[0].Attributes.MenuItem)
			return nil
		},
	}

}
