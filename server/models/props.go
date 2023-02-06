package models

import (
	"time"
)

// these structs are meant for direct conversion to and from json
type MovieProps struct {
	ID        uint
	Title     string
	Year      uint
	Director  []string
	PosterURL string
	Genre     []string
}

type WatchProps struct {
	MovieID   uint
	Date      time.Time
	Rating    Rating
	UserGenre []string
}

func PropsToMovie(props MovieProps) (m Movie) {
	m.ID = props.ID
	m.Title = props.Title
	m.Year = props.Year
	m.Director = props.Director
	m.PosterURL = props.PosterURL

	genres, themes, settings, tags := ConvertGenreStrings(props.Genre)
	m.Genre = genres
	m.Theme = themes
	m.Setting = settings
	m.Tag = tags

	return
}

func PropsToWatch(props WatchProps) (w Watch) {
	w.MovieID = props.MovieID
	w.Date = props.Date
	w.Rating = props.Rating

	genres, themes, settings, tags := ConvertGenreStrings(props.UserGenre)
	w.Genre = genres
	w.Theme = themes
	w.Setting = settings
	w.Tag = tags

	return
}
