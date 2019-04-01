package lunch

import (
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/minhajuddinkhan/cogs/types"
)

// Today gets todays lunch
func Today(store bolt.Store, creds *types.Credentials) (*types.Lunch, error) {

	raw, err := Request(store, creds)
	if err != nil {
		return nil, err
	}
	return Format(raw)

}
