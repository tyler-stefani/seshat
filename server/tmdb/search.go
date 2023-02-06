package tmdb

import (
	"os"
	"seshat-server/models"
	"strconv"

	tmdb "github.com/cyruzin/golang-tmdb"
)

type TMDB struct {
	client *tmdb.Client
}

func (t *TMDB) Connect() {
	if client, err := tmdb.Init(os.Getenv("TMDB_API_KEY")); err != nil {
		panic(err)
	} else {
		t.client = client
	}
}

func (t *TMDB) Search(title string, year string) []models.MovieProps {
	results := []models.MovieProps{}
	options := make(map[string]string)
	options["year"] = year
	options["region"] = "US"
	searchResults, err := t.client.GetSearchMovies(title, options)
	if err != nil {
		panic(err)
	}

	for _, searchResult := range searchResults.Results {
		id := int(searchResult.ID)

		title := searchResult.Title
		if searchResult.Title != searchResult.OriginalTitle {
			title = title + " (" + searchResult.OriginalTitle + ")"
		}

		year, _ := strconv.Atoi(searchResult.ReleaseDate[:4])

		credits, _ := t.client.GetMovieCredits(id, nil)
		directors := []string{}
		for _, person := range credits.Crew {
			if person.Job == "Director" {
				directors = append(directors, person.Name)
			}
		}

		genres := []string{}
		details, _ := t.client.GetMovieDetails(id, nil)
		for _, genre := range details.Genres {
			genres = append(genres, genre.Name)
		}

		allKeywords, _ := t.client.GetMovieKeywords(id)
		for _, keyword := range allKeywords.Keywords {
			switch keyword.Name {
			case "anthology":
				genres = append(genres, "Anthology")
			case "biography":
				genres = append(genres, "Biopic")
			case "coming of age":
				genres = append(genres, "Coming of Age")
			case "dark comedy":
				genres = append(genres, "Dark Comedy")
			case "slasher":
				genres = append(genres, "Slasher")
			case "superhero":
				genres = append(genres, "Superhero")
			}
		}

		var currMovie = &models.MovieProps{
			ID:        uint(id),
			Title:     title,
			Year:      uint(year),
			Director:  directors,
			PosterURL: tmdb.GetImageURL(searchResult.PosterPath, tmdb.Original),
			Genre:     genres,
		}

		results = append(results, *currMovie)
	}
	return results
}
