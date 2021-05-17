package anime

import (
	"encoding/json"
	"fmt"
)

//Sort is an enum reprenting the different sorting possibilities.
type Sort int

const (
	//SortByScore sorts by score (from highest score to lowest)
	SortByScore Sort = iota + 1
	//SortByUserListCount sorts by how many users have listed the anime (from most listed to least)
	SortByUserListCount
)

var sortStrDict = map[Sort]string{
	SortByScore:         "anime_score",
	SortByUserListCount: "anime_num_list_users",
}

//String returns the string representation of an enum value.
//If the value is not valid (e.g. undefined enum value), this functions returns "undefined".
func (sort Sort) String() string {
	sortStr := "undefined"

	if str, ok := sortStrDict[sort]; ok {
		sortStr = str
	}

	return sortStr
}

func (s *Sort) UnmarshalJSON(b []byte) error {
	var sortStr string

	if err := json.Unmarshal(b, &sortStr); err != nil {
		return err
	}

	found := false

	for k, v := range sortStrDict {
		if v == sortStr {
			*s = k
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%q is not a sorting type", sortStr)
	}

	return nil
}
