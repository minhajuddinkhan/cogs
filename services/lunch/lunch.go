package lunch

import (
	"github.com/minhajuddinkhan/cogs/services/cogs"
	"github.com/minhajuddinkhan/cogs/types"
)

// Today gets todays lunch
func Today(c cogs.Cogs) (*types.Lunch, error) {

	raw, err := Request(c)
	if err != nil {
		return nil, err
	}
	return Format(raw)

}
