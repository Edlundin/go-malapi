package anime

import "time"

//ListStatus represents a MyAnimeList list.
type ListStatus int

const (
	//ListStatusWatching represents the "watching" list.
	ListStatusWatching ListStatus = iota + 1
	//ListStatusCompleted represents the "completed" list.
	ListStatusCompleted
	//ListStatusOnHold represents the "on hold" list.
	ListStatusOnHold
	//ListStatusDropped represents the "dropped" list.
	ListStatusDropped
	//ListStatusPlanToWatch represents the "plan to watch" list.
	ListStatusPlanToWatch
)

var listStatusStrDict = map[ListStatus]string{
	ListStatusWatching:    "watching",
	ListStatusCompleted:   "completed",
	ListStatusOnHold:      "on_hold",
	ListStatusDropped:     "dropped",
	ListStatusPlanToWatch: "plan_to_watch",
}

func (season ListStatus) String() string {
	listStatusStr := "unknown"

	if str, ok := listStatusStrDict[season]; ok {
		listStatusStr = str
	}

	return listStatusStr
}

//ListStatusObject represents a MyListStatus JSON object
type ListStatusObject struct {
	Comments            string     `json:"comments"`
	IsRewatching        bool       `json:"is_rewatching"`
	WatchedEpisodeCount int        `json:"num_episodes_watched"`
	RewatchedCount      int        `json:"num_times_rewatched"`
	Priority            int        `json:"priority"`
	RewatchedValue      int        `json:"rewatch_value"`
	Score               float64    `json:"score"`
	Status              ListStatus `json:"status"`
	Tags                []string   `json:"tags"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
