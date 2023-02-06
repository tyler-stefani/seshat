package models

import "time"

type Score uint

const (
	scoreZero Score = iota
	scoreOne
	scoreTwo
	scoreThree
	scoreFour
	scoreFive
)

type Rating struct {
	Enjoyment, Feel, Style, Characters, Narrative, Impact, Interest Score
}

type Movie struct {
	ID        uint
	Title     string
	Year      uint
	Director  []string
	PosterURL string
	Genre     Genres
	Theme     Themes
	Setting   Settings
	Tag       Tags
}

type Watch struct {
	MovieID uint
	Date    time.Time
	Rating  Rating
	Genre   Genres
	Theme   Themes
	Setting Settings
	Tag     Tags
}
