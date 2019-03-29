package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/minhajuddinkhan/cogs"
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
func Lunch(store bolt.Store, creds *types.Credentials) cli.Command {
	return cli.Command{
		Name:  "lunch",
		Usage: "Gets you todays lunch",
		Action: func(c *cli.Context) error {
			logrus.Info("fetching lunch..")
			raw, err := cogs.Lunch(creds.AccessToken)
			if err != nil {
				if err := cogs.Update(store, creds); err != nil {
					return err
				}
			}
			var lb lunchBody
			if err := json.Unmarshal(raw, &lb); err != nil {
				return err
			}
			fmt.Printf("Todays lunch is %s \n", lb.Data[0].Attributes.MenuItem)
			return nil
		},
	}

}
