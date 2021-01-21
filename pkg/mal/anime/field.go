package anime

type AnimeField string

const (
	FieldID                               AnimeField = "id"                                   //integer
	FieldTitle                            AnimeField = "title"                                //string
	FieldMainPictureMedium                AnimeField = "main_picture{medium}"                 //string
	FieldMainPictureLarge                 AnimeField = "main_picture{large}"                  //string
	FieldAlternativeTitlesSynonyms        AnimeField = "alternative_titles{synonyms}"         //[]string
	FieldAlternativeTitlesEn              AnimeField = "alternative_titles{en}"               //string
	FieldAlternativeTitlesJa              AnimeField = "alternative_titles{ja}"               //string
	FieldStartDate                        AnimeField = "start_date"                           //YYYY-MM-DD
	FieldEndDate                          AnimeField = "end_date"                             //YYYY-MM-DD
	FieldSynopsis                         AnimeField = "synopsis"                             //string
	FieldMeanScore                        AnimeField = "mean"                                 //float
	FieldRank                             AnimeField = "rank"                                 //integer
	FieldPopularity                       AnimeField = "popularity"                           //integer
	FieldEnlisting                        AnimeField = "num_list_users"                       //integer
	FieldScoringUserCount                 AnimeField = "num_scoring_users"                    //integer
	FieldNsfw                             AnimeField = "nsfw"                                 //bool
	FieldCreatedAt                        AnimeField = "created_at"                           //string i.e. 2015-03-02T06:03:11+00:00 (RFC 3339, ISO 8601)
	FieldUpdatedAt                        AnimeField = "updated_at"                           //string i.e. 2018-04-25T09:14:14+00:00 (RFC 3339, ISO 8601)
	FieldMediaType                        AnimeField = "media_type"                           //string //TODO: make enum MediaType
	FieldStatus                           AnimeField = "finished_airing"                      //string //TODO: make enum AnimeStatus
	FieldGenres                           AnimeField = "genres"                               //[]string //TODO: create object ([]AnimeGenre)
	FieldMyListStatusStatus               AnimeField = "my_list_status{status}"               //string i.e. plan_to_watch, watched, ... //TODO: make enum of AnimeList
	FieldMyListStatusScore                AnimeField = "my_list_status{score}"                //integer
	FieldMyListStatusWatchedEpisodesCount AnimeField = "my_list_status{num_episodes_watched}" //integer
	FieldMyListStatusIsRewatching         AnimeField = "my_list_status{is_rewatching}"        //bool
	FieldMyListStatusUpdatedAt            AnimeField = "my_list_status{updated_at}"           //string i.e. 2017-11-11T19:51:22+00:00 (RFC 3339, ISO 8601)
	FieldEpisodeCount                     AnimeField = "num_episodes"                         //integer
	FieldStartSeasonYear                  AnimeField = "start_season{year}"                   //integer
	FieldStartSeasonSeason                AnimeField = "start_season{season}"                 //AnimeSeason string values
	FieldBroadcastDayOfTheWeek            AnimeField = "broadcast{day_of_the_week}"           //string //TODO: make enum of DayOfTheWeek
	FieldBroadcastStartTime               AnimeField = "broadcast{start_time}"                //string i.e. 18:00
	FieldSource                           AnimeField = "source"                               //string //TODO: make enum of AnimeSourceMaterial
	FieldAverageEpisodeDuration           AnimeField = "average_episode_duration"             //integer
	FieldPgRating                         AnimeField = "rating"                               //string //TODO: make enum of AnimePgRating
	FieldPictures                         AnimeField = "pictures"                             //[]string //TODO: create object ([]AnimePicture) //TODO: if key different of "medium" or "large" print warning
	FieldBackground                       AnimeField = "background"                           //string
	FieldRelatedAnime                     AnimeField = "related_anime"                        //[]string //TODO: create object ([]RelatedAnime)
	FieldRelatedManga                     AnimeField = "related_manga"                        //[]string //TODO: create object ([]RelatedManga)
	FieldRecommendations                  AnimeField = "recommendations"                      //[]string //TODO: create object ([]AnimeRecommendation)
	FieldStudios                          AnimeField = "studios"                              //[]string //TODO: create object ([]AnimeStudio or []Studio)
	FieldStatisticsStatus                 AnimeField = "statistics{status}"                   //[]string //TODO: create object ([]AnimeStudio or []Studio)
	FieldStatisticsEnlistedCount          AnimeField = "statistics{num_list_users}"           //[]string //TODO: create object ([]AnimeStudio or []Studio)
)
