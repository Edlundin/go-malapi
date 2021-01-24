package anime

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Edlundin/go-malapi/pkg/mal"
)

func animeFieldsToGetParams(animeFields []AnimeField) string { //! TODO: test this function
	var getParams string

	if len(animeFields) > 0 {
		sort.SliceStable(animeFields, func(i int, j int) bool {
			return animeFields[i] < animeFields[j]
		})

		getParams = fmt.Sprintf("fields=%s", animeFields[0])

		for i := 1; i < len(animeFields); i++ {
			const subtypeDelimiter string = "{"
			previousField, currentField := string(animeFields[i-1]), string(animeFields[i])
			previousIndex, currentIndex := strings.Index(previousField, subtypeDelimiter), strings.Index(currentField, subtypeDelimiter)

			if currentIndex > -1 && previousIndex > -1 && currentField[:currentIndex] == previousField[:previousIndex] { //previousField and currentField are two subtypes from the same category
				getParams = fmt.Sprintf("%s,%s", strings.TrimSuffix(getParams, "}"), currentField[currentIndex+1:]) //TODO: example of what it does
			} else { //previousField and currentField are two subtypes from different categories, or one of them is not a subtype
				getParams += fmt.Sprintf(",%s", currentField)
			}
		}
	}

	return getParams
}

//if resultLimit == 0 -> 100
//offset default: 0
func SearchAnime(apiClient mal.Client, animeTitle string, resultLimit uint, offset uint, fields []AnimeField) {
	const endPoint string = "/anime"
}

func GetAnimeDetails(apiClient mal.Client, animeId uint, fields []AnimeField) {
	var endPoint string = fmt.Sprintf("/anime/%d", animeId)
}

//if resultLimit == 0 -> 100
//offset default: 0
func GetAnimeRanking(apiClient mal.Client, rankingType AnimeRankingType, resultLimit uint, offset uint, fields []AnimeField) {
	const endPoint string = "/anime/ranking"
}

func GetSeasonalAnimes(apiClient mal.Client, year uint, season AnimeSeason) {
	var endPoint string = fmt.Sprintf("/anime/ranking/%d/%s", year, season)
}

//if resultLimit == 0 -> 100
//offset default: 0
func GetSuggestedAnimes(apiClient mal.Client, resultLimit uint, offset uint) {
	const endPoint string = "/anime/suggestion"
}
