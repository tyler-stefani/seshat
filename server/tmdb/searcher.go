package tmdb

import ("seshat-server/models")

type Searcher interface{
	Search(string, string) []models.MovieProps
}