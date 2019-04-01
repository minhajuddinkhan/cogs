package lunch

import (
	"fmt"
	"net/http"

	"github.com/minhajuddinkhan/cogs/endpoint"
	"github.com/minhajuddinkhan/cogs/services/cogs"
)

const (
	url = endpoint.BaseURL + "/Lunches/Weekly"
)

// Request requests lunch from cogs API
func Request(c cogs.Cogs) ([]byte, error) {

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.Credentials().AccessToken),
	}
	raw, err := endpoint.Request(url, http.MethodGet, "", headers)
	if err != nil {
		switch e := err.(type) {
		case endpoint.HttpError:
			if e.Code != http.StatusUnauthorized {
				return nil, e
			}
			if err := c.Update(); err != nil {
				return nil, err
			}
		default:
			return nil, err
		}
	}
	return raw, nil
}
