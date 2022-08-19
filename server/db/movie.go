package db

type Rating struct {
	Enjoyment, Feel, Style, Characters, Narrative, Impact, Interest Number
}

type Movie struct {
	Id        RichText
	Title     Title
	Year      Number
	Poster    File
	Directors MultiSelect
	Rating    Rating
	Genre     MultiSelect
	Watched   Date
}
