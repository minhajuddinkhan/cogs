package types

import (
	"strings"
	"time"
)

type LunchRequestBody struct {
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

type Lunch struct {
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
