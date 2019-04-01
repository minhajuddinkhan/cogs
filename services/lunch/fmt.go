package lunch

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/minhajuddinkhan/cogs/types"
)

func Format(raw []byte) (*types.Lunch, error) {

	var lb types.LunchRequestBody
	if err := json.Unmarshal(raw, &lb); err != nil {
		return nil, err
	}
	var out string
	for _, l := range lb.Data {
		if l.Attributes.Date.Format("2006-01-02") == time.Now().Format("2006-01-02") {
			out = l.Attributes.MenuItem
		}
	}
	if out == "" {
		return nil, fmt.Errorf("unable to find lunch today")
	}

	var food types.Lunch
	isReg := true
	for j, x := range strings.Split(out, ",") {
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
	return &food, nil
}
