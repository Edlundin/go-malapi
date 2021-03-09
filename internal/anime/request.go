package anime

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/Edlundin/go-malapi/pkg/mal/anime"
)

func animeFieldsToQueryArguments(animeFields []anime.Field) string {
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

//GenerateGetAnimeListRequest creates and returns a request. Use it to obtain a list of animes based on a title.
//The base URL is the common part of all the API's endpoints.
//The bearer token is the access token obtained through the OAuth2 identification.
//The anime title is what an anime is search against.
//The limit parameter limits the maximum number of animes returned by the API. If limit equals 0, MAL's API will return a list of up to a hundred animes (a hundred is the default value).
//The offset (or paging) is used to ignore the first x animes found by the API (where x is the offset value). The ignored animes are not counted toward the limit. The default offset value is 0.
//The fields limits the details per anime returned by the API. If the array is empty, the API will return all the available fields for all found animes.
func GenerateGetAnimeListRequest(malAPIBaseURL string, bearerToken string, animeTitle string, limit uint, offset uint, fields []anime.Field) (*http.Request, error) {
	const endPoint string = "/anime"

	if limit == 0 {
		limit = 100
	}

	req, err := http.NewRequest(http.MethodGet, malAPIBaseURL+endPoint, nil)

	if err != nil {
		return nil, fmt.Errorf("generating a search anime request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	queryArguments := req.URL.Query()
	queryArguments.Add("q", animeTitle)
	queryArguments.Add("limit", strconv.FormatUint(uint64(limit), 10))
	queryArguments.Add("offset", strconv.FormatUint(uint64(offset), 10))
	queryArguments.Add("fields", animeFieldsToQueryArguments(fields))

	req.URL.RawQuery = queryArguments.Encode()

	return req, nil
}

//GenerateGetAnimeDetailsRequest creates and returns a request. Use it to obtain the details of a particular anime.
//The base URL is the common part of all the API's endpoints.
//The bearer token is the access token obtained through the OAuth2 identification.
//The anime ID of the anime the API will get details for.
//The fields limits the details per anime returned by the API. If the array is empty, the API will return all the available fields for all found animes.
func GenerateGetAnimeDetailsRequest(malAPIBaseURL string, bearerToken string, animeID uint, fields []anime.Field) (*http.Request, error) {
	var endPoint string = fmt.Sprintf("/anime/%d", animeID)

	req, err := http.NewRequest(http.MethodGet, malAPIBaseURL+endPoint, nil)

	if err != nil {
		return nil, fmt.Errorf("generating a anime details request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	queryArguments := req.URL.Query()
	queryArguments.Add("fields", animeFieldsToQueryArguments(fields))

	req.URL.RawQuery = queryArguments.Encode()

	return req, nil
}

//GenerateGetAnimeRankingRequest creates and returns a request. Use it to obtain the animes and their rank in a ranking list (e.g. the "airing" ranking type returns the top airing animes).
//The base URL is the common part of all the API's endpoints.
//The bearer token is the access token obtained through the OAuth2 identification.
//The ranking type from which ranking list the API will return the animes.
//The limit parameter limits the maximum number of animes returned by the API. If limit equals 0, MAL's API will return a list of up to a hundred animes (a hundred is the default value).
//The offset (or paging) is used to ignore the first x animes found by the API (where x is the offset value). The ignored animes are not counted toward the limit. The default offset value is 0.
//The fields limits the details per anime returned by the API. If the array is empty, the API will return all the available fields for all found animes.
func GenerateGetAnimeRankingRequest(malAPIBaseURL string, bearerToken string, rankingType anime.RankingType, limit uint, offset uint, fields []anime.Field) (*http.Request, error) {
	const endPoint string = "/anime/ranking"

	if limit == 0 {
		limit = 100
	}

	req, err := http.NewRequest(http.MethodGet, malAPIBaseURL+endPoint, nil)

	if err != nil {
		return nil, fmt.Errorf("generating an anime ranking request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	queryArguments := req.URL.Query()
	queryArguments.Add("ranking_type", rankingType.String())
	queryArguments.Add("limit", strconv.FormatUint(uint64(limit), 10))
	queryArguments.Add("offset", strconv.FormatUint(uint64(offset), 10))
	queryArguments.Add("fields", animeFieldsToQueryArguments(fields))

	req.URL.RawQuery = queryArguments.Encode()

	return req, nil
}

//GenerateGetSeasonalAnimes creates and returns a request. Use it to obtain the list of animes airing for a particular season.
//The base URL is the common part of all the API's endpoints.
//The bearer token is the access token obtained through the OAuth2 identification.
//The limit parameter limits the maximum number of animes returned by the API. If limit equals 0, MAL's API will return a list of up to a hundred animes (a hundred is the default value).
//The offset (or paging) is used to ignore the first x animes found by the API (where x is the offset value). The ignored animes are not counted toward the limit. The default offset value is 0.
//The fields limits the details per anime returned by the API. If the array is empty, the API will return all the available fields for all found animes.
func GenerateGetSeasonalAnimes(malAPIBaseURL string, bearerToken string, year uint, season anime.Season, sortBy anime.Sort, limit uint, offset uint, fields []anime.Field) (*http.Request, error) {
	var endPoint string = fmt.Sprintf("/anime/ranking/%d/%s", year, season)

	if limit == 0 {
		limit = 100
	}

	req, err := http.NewRequest(http.MethodGet, malAPIBaseURL+endPoint, nil)

	if err != nil {
		return nil, fmt.Errorf("generating a seasonal animes request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	queryArguments := req.URL.Query()
	queryArguments.Add("sort", sortBy.String())
	queryArguments.Add("limit", strconv.FormatUint(uint64(limit), 10))
	queryArguments.Add("offset", strconv.FormatUint(uint64(offset), 10))
	queryArguments.Add("fields", animeFieldsToQueryArguments(fields))

	req.URL.RawQuery = queryArguments.Encode()

	return req, nil
}

//GenerateGetSuggestedAnimes creates and returns a request. Use it to obtain the suggested animes of the account linked to the bearer token.
//The base URL is the common part of all the API's endpoints.
//The bearer token is the access token obtained through the OAuth2 identification.
//The limit parameter limits the maximum number of animes returned by the API. If limit equals 0, MAL's API will return a list of up to a hundred animes (a hundred is the default value).
//The offset (or paging) is used to ignore the first x animes found by the API (where x is the offset value). The ignored animes are not counted toward the limit. The default offset value is 0.
//The fields limits the details per anime returned by the API. If the array is empty, the API will return all the available fields for all found animes.
func GenerateGetSuggestedAnimes(malAPIBaseURL string, bearerToken string, limit uint, offset uint, fields []anime.Field) (*http.Request, error) {
	const endPoint string = "/anime/suggestion"

	if limit == 0 {
		limit = 100
	}

	req, err := http.NewRequest(http.MethodGet, malAPIBaseURL+endPoint, nil)

	if err != nil {
		return nil, fmt.Errorf("generating a suggested animes request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	queryArguments := req.URL.Query()
	queryArguments.Add("limit", strconv.FormatUint(uint64(limit), 10))
	queryArguments.Add("offset", strconv.FormatUint(uint64(offset), 10))
	queryArguments.Add("fields", animeFieldsToQueryArguments(fields))

	req.URL.RawQuery = queryArguments.Encode()

	return req, nil
}
