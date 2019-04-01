package lunch

import (
	"fmt"
	"net/http"

	"github.com/minhajuddinkhan/cogs"
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/minhajuddinkhan/cogs/types"
)

const (
	url = "/Lunches/Weekly"
)

// Request requests lunch from cogs API
func Request(store bolt.Store, creds *types.Credentials) ([]byte, error) {

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", creds.AccessToken),
	}
	raw, err := cogs.Request(url, http.MethodGet, "", headers)
	if err != nil {
		switch e := err.(type) {
		case cogs.HttpError:
			if e.Code != http.StatusUnauthorized {
				return nil, e
			}
			if err := cogs.Update(store, creds); err != nil {
				return nil, err
			}
		default:
			return nil, err
		}
	}
	return raw, nil
}
