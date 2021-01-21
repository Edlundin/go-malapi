package anime

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Edlundin/go-malapi/pkg/mal"
)

//if resultLimit == 0 -> 100
//offset default: 0
func SearchAnime(apiClient mal.Client, animeTitle string, resultLimit uint, offset uint, fields []AnimeField) {
}

func GetAnimeDetails(apiClient mal.Client, animeId uint, fields []AnimeField) {}

//if resultLimit == 0 -> 100
//offset default: 0
func GetAnimeRanking(apiClient mal.Client, rankingType AnimeRankingType, resultLimit uint, offset uint, fields []AnimeField) {
}

func GetSeasonalAnimes(apiClient mal.Client, year uint, season AnimeSeason) {}

//if resultLimit == 0 -> 100
//offset default: 0
func GetSuggestedAnimes(apiClient mal.Client, resultLimit uint, offset uint) {}
