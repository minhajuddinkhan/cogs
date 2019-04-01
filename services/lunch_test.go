package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/minhajuddinkhan/cogs/services/lunch"
	"github.com/stretchr/testify/assert"
)

func TestLunchFmt(t *testing.T) {
	date := time.Now().Format("2006-01-02T15:04:05")
	lunchName := "Regular Lunch, Daal Chawal"
	b := fmt.Sprintf(`{"data":[{"attributes": {"lunch-date":"%s","menu-item": "%s"}}]}`, date, lunchName)
	raw := []byte(b)
	lunchToday, err := lunch.Format(raw)
	assert.Nil(t, err)
	assert.Equal(t, lunchToday.Regular[0], "Daal Chawal")
}
