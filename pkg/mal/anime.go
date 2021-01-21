package mal

import (
	"fmt"
	"sort"
	"strings"
)

type AnimeField string

const (
	AnimeField_ID                                    AnimeField = "id"                                   //integer
	AnimeField_TITLE                                 AnimeField = "title"                                //string
	AnimeField_MAIN_PICTURE_MEDIUM                   AnimeField = "main_picture{medium}"                 //string
	AnimeField_MAIN_PICTURE_LARGE                    AnimeField = "main_picture{large}"                  //string
	AnimeField_ALTERNATIVE_TITLES_SYNONYMS           AnimeField = "alternative_titles{synonyms}"         // []string
	AnimeField_ALTERNATIVE_TITLES_EN                 AnimeField = "alternative_titles{en}"               //string
	AnimeField_ALTERNATIVE_TITLES_JA                 AnimeField = "alternative_titles{ja}"               //string
	AnimeField_START_DATE                            AnimeField = "start_date"                           //YYYY-MM-DD
	AnimeField_END_DATE                              AnimeField = "end_date"                             //YYYY-MM-DD
	AnimeField_SYNOPSIS                              AnimeField = "synopsis"                             //string
	AnimeField_MEAN_SCORE                            AnimeField = "mean"                                 //float
	AnimeField_RANK                                  AnimeField = "rank"                                 //integer
	AnimeField_POPULARITY                            AnimeField = "popularity"                           //integer
	AnimeField_ENLISTING_COUNT                       AnimeField = "num_list_users"                       //integer
	AnimeField_SCORING_USER_COUNT                    AnimeField = "num_scoring_users"                    //integer
	AnimeField_NSFW                                  AnimeField = "nsfw"                                 //bool
	AnimeField_CREATED_AT                            AnimeField = "created_at"                           //string i.e. 2015-03-02T06:03:11+00:00 (RFC 3339, ISO 8601)
	AnimeField_UPDATED_AT                            AnimeField = "updated_at"                           //string i.e. 2018-04-25T09:14:14+00:00 (RFC 3339, ISO 8601)
	AnimeField_MEDIA_TYPE                            AnimeField = "media_type"                           //string //TODO: make enum MediaType
	AnimeField_STATUS                                AnimeField = "finished_airing"                      //string //TODO: make enum AnimeStatus
	AnimeField_GENRES                                AnimeField = "genres"                               //[]string //TODO: create object ([]AnimeGenre)
	AnimeField_MY_LIST_STATUS_STATUS                 AnimeField = "my_list_status{status}"               //string i.e. plan_to_watch, watched, ... //TODO: make enum of AnimeList
	AnimeField_MY_LIST_STATUS_SCORE                  AnimeField = "my_list_status{score}"                //integer
	AnimeField_MY_LIST_STATUS_WATCHED_EPISODES_COUNT AnimeField = "my_list_status{num_episodes_watched}" //integer
	AnimeField_MY_LIST_STATUS_IS_REWATCHING          AnimeField = "my_list_status{is_rewatching}"        //bool
	AnimeField_MY_LIST_STATUS_UPDATED_AT             AnimeField = "my_list_status{updated_at}"           //string i.e. 2017-11-11T19:51:22+00:00 (RFC 3339, ISO 8601)
	AnimeField_EPISODE_COUNT                         AnimeField = "num_episodes"                         //integer
	AnimeField_START_SEASON_YEAR                     AnimeField = "start_season{year}"                   //integer
	AnimeField_START_SEASON_SEASON                   AnimeField = "start_season{season}"                 //AnimeSeason string values
	AnimeField_BROADCAST_DAY_OF_THE_WEEK             AnimeField = "broadcast{day_of_the_week}"           //string //TODO: make enum of DayOfTheWeek
	AnimeField_BROADCAST_START_TIME                  AnimeField = "broadcast{start_time}"                //string i.e. 18:00
	AnimeField_SOURCE                                AnimeField = "source"                               //string //TODO: make enum of AnimeSourceMaterial
	AnimeField_AVERAGE_EPISODE_DURATION              AnimeField = "average_episode_duration"             //integer
	AnimeField_PG_RATING                             AnimeField = "rating"                               //string //TODO: make enum of AnimePgRating
	AnimeField_PICTURES                              AnimeField = "pictures"                             //[]string //TODO: create object ([]AnimePicture) //TODO: if key different of "medium" or "large" print warning
	AnimeField_BACKGROUND                            AnimeField = "background"                           //string
	AnimeField_RELATED_ANIME                         AnimeField = "related_anime"                        //[]string //TODO: create object ([]RelatedAnime)
	AnimeField_RELATED_MANGA                         AnimeField = "related_manga"                        //[]string //TODO: create object ([]RelatedManga)
	AnimeField_RECOMMENDATIONS                       AnimeField = "recommendations"                      //[]string //TODO: create object ([]AnimeRecommendation)
	AnimeField_STUDIOS                               AnimeField = "studios"                              //[]string //TODO: create object ([]AnimeStudio or []Studio)
	AnimeField_STATISTICS_STATUS                     AnimeField = "statistics{status}"                   //[]string //TODO: create object ([]AnimeStudio or []Studio)
	AnimeField_STATISTICS_ENLISTED_COUNT             AnimeField = "statistics{num_list_users}"           //[]string //TODO: create object ([]AnimeStudio or []Studio)
)

type AnimeRankingType string

const (
	AnimeRankingType_ALL        AnimeRankingType = "all"
	AnimeRankingType_AIRING     AnimeRankingType = "airing"
	AnimeRankingType_UPCOMING   AnimeRankingType = "upcoming"
	AnimeRankingType_TV         AnimeRankingType = "tv"
	AnimeRankingType_OVA        AnimeRankingType = "ova"
	AnimeRankingType_MOVIE      AnimeRankingType = "movie"
	AnimeRankingType_SPECIAL    AnimeRankingType = "special"
	AnimeRankingType_POPULARITY AnimeRankingType = "bypopularity"
	AnimeRankingType_FAVORITE   AnimeRankingType = "favorite"
	AnimeRankingType_UNKNOWN    AnimeRankingType = "unknown" //never used by MyAnimeList. Only for error handling.
)

type AnimeSeason string

const (
	AnimeSeason_WINTER  AnimeSeason      = "winter"
	AnimeSeason_SPRING  AnimeSeason      = "spring"
	AnimeSeason_SUMMER  AnimeSeason      = "summer"
	AnimeSeason_FALL    AnimeSeason      = "fall"
	AnimeSeason_UNKNOWN AnimeRankingType = "unknown" //never used by MyAnimeList. Only for error handling.
)

//if resultLimit == 0 -> 100
//offset default: 0
func SearchAnime(apiClient Client, animeTitle string, resultLimit uint, offset uint, fields []AnimeField) {
}

func GetAnimeDetails(apiClient Client, animeId uint, fields []AnimeField) {}

//if resultLimit == 0 -> 100
//offset default: 0
func GetAnimeRanking(apiClient Client, rankingType AnimeRankingType, resultLimit uint, offset uint, fields []AnimeField) {
}

func GetSeasonalAnimes(apiClient Client, year uint, season AnimeSeason) {}

//if resultLimit == 0 -> 100
//offset default: 0
func GetSuggestedAnimes(apiClient Client, resultLimit uint, offset uint) {}
