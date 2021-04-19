package anime

type Field string

const (
	FieldID                               Field = "id"
	FieldTitle                            Field = "title"
	FieldMainPictureMedium                Field = "main_picture{medium}"
	FieldMainPictureLarge                 Field = "main_picture{large}"
	FieldAlternativeTitlesSynonyms        Field = "alternative_titles{synonyms}"
	FieldAlternativeTitlesEn              Field = "alternative_titles{en}"
	FieldAlternativeTitlesJa              Field = "alternative_titles{ja}"
	FieldStartDate                        Field = "start_date"
	FieldEndDate                          Field = "end_date"
	FieldSynopsis                         Field = "synopsis"
	FieldMeanScore                        Field = "mean"
	FieldRank                             Field = "rank"
	FieldPopularity                       Field = "popularity"
	FieldEnlisting                        Field = "num_list_users"
	FieldScoringUserCount                 Field = "num_scoring_users"
	FieldNsfw                             Field = "nsfw"
	FieldCreatedAt                        Field = "created_at"
	FieldUpdatedAt                        Field = "updated_at"
	FieldMediaType                        Field = "media_type"
	FieldStatus                           Field = "finished_airing" //* only works when submitted to a particular anime (/anime/$id)
	FieldGenres                           Field = "genres"
	FieldMyListStatusStatus               Field = "my_list_status{status}"
	FieldMyListStatusScore                Field = "my_list_status{score}"
	FieldMyListStatusWatchedEpisodesCount Field = "my_list_status{num_episodes_watched}"
	FieldMyListStatusIsRewatching         Field = "my_list_status{is_rewatching}"
	FieldMyListStatusUpdatedAt            Field = "my_list_status{updated_at}"
	FieldEpisodeCount                     Field = "num_episodes"
	FieldStartSeasonYear                  Field = "start_season{year}"
	FieldStartSeasonSeason                Field = "start_season{season}"
	FieldBroadcastDayOfTheWeek            Field = "broadcast{day_of_the_week}"
	FieldBroadcastStartTime               Field = "broadcast{start_time}"
	FieldSource                           Field = "source"
	FieldAverageEpisodeDuration           Field = "average_episode_duration"
	FieldPgRating                         Field = "rating"
	FieldPictures                         Field = "pictures"        //* only works when submitted to a particular anime (/anime/$id)
	FieldBackground                       Field = "background"      //* only works when submitted to a particular anime (/anime/$id)
	FieldRelatedAnime                     Field = "related_anime"   //* only works when submitted to a particular anime (/anime/$id)
	FieldRelatedManga                     Field = "related_manga"   //* only works when submitted to a particular anime (/anime/$id)
	FieldRecommendations                  Field = "recommendations" //* only works when submitted to a particular anime (/anime/$id), returns []Anime
	FieldStudios                          Field = "studios"
	FieldStatisticsStatus                 Field = "statistics{status}"
	FieldStatisticsEnlistedCount          Field = "statistics{num_list_users}"
)
