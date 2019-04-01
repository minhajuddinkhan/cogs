package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/minhajuddinkhan/cogs/services/lunch"
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/minhajuddinkhan/cogs/types"
	"github.com/urfave/cli"
)

// Lunch Gets Lunch
func Lunch(store bolt.Store, creds *types.Credentials) cli.Command {
	return cli.Command{
		Name:  "lunch",
		Usage: "Gets you todays lunch",
		Action: func(c *cli.Context) error {
			logrus.Info("fetching lunch..")

			todaysLunch, err := lunch.Today(store, creds)
			if err != nil {
				return err
			}

			color.White("Regular Lunch")
			for _, l := range todaysLunch.Regular {
				fmt.Println(l)
			}
			color.Green("Low Calorie")
			for _, l := range todaysLunch.LowCal {
				fmt.Println(l)
			}

			return nil

		},
	}
}
