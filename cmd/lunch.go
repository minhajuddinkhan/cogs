package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/fatih/color"
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
			Date     Date   `json:"lunch-date"`
			Likes    int    `json:"likes-count"`
			Dislikes int    `json:"dislikes-count"`
		} `json:"attributes"`
	} `json:"data"`
}

type lunch struct {
	LowCal  []string
	Regular []string
}

const splitter = ","

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	str := strings.TrimSpace(string(b))
	str = str[1 : len(str)-1]
	if str == "" {
		d.Time = time.Time{}
		return nil
	}

	t, err := time.Parse("2006-01-02T15:04:05", str)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
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
			var out string
			for _, l := range lb.Data {
				if l.Attributes.Date.Format("2006-01-02") == time.Now().Format("2006-01-02") {
					out = l.Attributes.MenuItem
				}
			}
			if out == "" {
				return fmt.Errorf("unable to find lunch today")
			}

			var food lunch
			isReg := true
			for j, x := range strings.Split(out, splitter) {
				x = strings.TrimSpace(x)
				if j == 0 {
					continue
				}
				if x == "Low Calories Lunch" {
					isReg = false
					continue
				}

				if isReg {
					food.Regular = append(food.Regular, x)
				} else {
					food.LowCal = append(food.LowCal, x)
				}

			}

			color.White("Regular lunch")
			for _, x := range food.Regular {
				fmt.Println(x)
			}
			color.White("Low Cal")
			for _, x := range food.LowCal {
				fmt.Println(x)
			}
			return nil
		},
	}

}
