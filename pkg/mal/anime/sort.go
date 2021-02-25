package anime

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

func (sort Sort) String() string {
	sortStr := "unknown"

	if str, ok := sortStrDict[sort]; ok {
		sortStr = str
	}

	return sortStr
}
