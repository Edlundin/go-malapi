package anime

type Field string

const (
	FieldID                               Field = "id"                                   //integer
	FieldTitle                            Field = "title"                                //string
	FieldMainPictureMedium                Field = "main_picture{medium}"                 //string
	FieldMainPictureLarge                 Field = "main_picture{large}"                  //string
	FieldAlternativeTitlesSynonyms        Field = "alternative_titles{synonyms}"         //[]string
	FieldAlternativeTitlesEn              Field = "alternative_titles{en}"               //string
	FieldAlternativeTitlesJa              Field = "alternative_titles{ja}"               //string
	FieldStartDate                        Field = "start_date"                           //YYYY-MM-DD
	FieldEndDate                          Field = "end_date"                             //YYYY-MM-DD
	FieldSynopsis                         Field = "synopsis"                             //string
	FieldMeanScore                        Field = "mean"                                 //float
	FieldRank                             Field = "rank"                                 //integer
	FieldPopularity                       Field = "popularity"                           //integer
	FieldEnlisting                        Field = "num_list_users"                       //integer
	FieldScoringUserCount                 Field = "num_scoring_users"                    //integer
	FieldNsfw                             Field = "nsfw"                                 //bool
	FieldCreatedAt                        Field = "created_at"                           //string i.e. 2015-03-02T06:03:11+00:00 (RFC 3339, ISO 8601)
	FieldUpdatedAt                        Field = "updated_at"                           //string i.e. 2018-04-25T09:14:14+00:00 (RFC 3339, ISO 8601)
	FieldMediaType                        Field = "media_type"                           //string //TODO: make enum MediaType
	FieldStatus                           Field = "finished_airing"                      //string //TODO: make enum AnimeStatus
	FieldGenres                           Field = "genres"                               //[]string //TODO: create object ([]AnimeGenre)
	FieldMyListStatusStatus               Field = "my_list_status{status}"               //string i.e. plan_to_watch, watched, ... //TODO: make enum of AnimeList
	FieldMyListStatusScore                Field = "my_list_status{score}"                //integer
	FieldMyListStatusWatchedEpisodesCount Field = "my_list_status{num_episodes_watched}" //integer
	FieldMyListStatusIsRewatching         Field = "my_list_status{is_rewatching}"        //bool
	FieldMyListStatusUpdatedAt            Field = "my_list_status{updated_at}"           //string i.e. 2017-11-11T19:51:22+00:00 (RFC 3339, ISO 8601)
	FieldEpisodeCount                     Field = "num_episodes"                         //integer
	FieldStartSeasonYear                  Field = "start_season{year}"                   //integer
	FieldStartSeasonSeason                Field = "start_season{season}"                 //AnimeSeason string values
	FieldBroadcastDayOfTheWeek            Field = "broadcast{day_of_the_week}"           //string //TODO: make enum of DayOfTheWeek
	FieldBroadcastStartTime               Field = "broadcast{start_time}"                //string i.e. 18:00
	FieldSource                           Field = "source"                               //string //TODO: make enum of AnimeSourceMaterial
	FieldAverageEpisodeDuration           Field = "average_episode_duration"             //integer
	FieldPgRating                         Field = "rating"                               //string //TODO: make enum of AnimePgRating
	FieldPictures                         Field = "pictures"                             //[]string //TODO: create object ([]AnimePicture) //TODO: if key different of "medium" or "large" print warning
	FieldBackground                       Field = "background"                           //string
	FieldRelatedAnime                     Field = "related_anime"                        //[]string //TODO: create object ([]RelatedAnime)
	FieldRelatedManga                     Field = "related_manga"                        //[]string //TODO: create object ([]RelatedManga)
	FieldRecommendations                  Field = "recommendations"                      //[]string //TODO: create object ([]AnimeRecommendation)
	FieldStudios                          Field = "studios"                              //[]string //TODO: create object ([]AnimeStudio or []Studio)
	FieldStatisticsStatus                 Field = "statistics{status}"                   //[]string //TODO: create object ([]AnimeStudio or []Studio)
	FieldStatisticsEnlistedCount          Field = "statistics{num_list_users}"           //[]string //TODO: create object ([]AnimeStudio or []Studio)
)
