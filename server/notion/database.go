package notion

import ("seshat-server/models")

type Writer interface {
	WriteMovie(models.Movie, string)
	LogWatch(models.Watch, string, string)
}