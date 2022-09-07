package main

import (
	"fmt"
	http "net/http"
	"os"
	"strconv"

	"seshat-server/db"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type basicMovie struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Year      int      `json:"year"`
	Directors []string `json:"directors"`
	Poster    string   `json:"poster"`
}

func search(c *gin.Context) {
	searchName := c.Query("name")
	searchYear := c.Query("year")

	results := []basicMovie{}

	tmdbClient, clientErr := tmdb.Init(os.Getenv("TMDB_API_KEY"))
	if clientErr != nil {
		fmt.Println(clientErr)
	}

	options := make(map[string]string)
	options["year"] = searchYear
	options["region"] = "US"
	searchResults, searchErr := tmdbClient.GetSearchMovies(searchName, options)
	if searchErr != nil {
		fmt.Println(searchErr)
	}

	for _, searchResult := range searchResults.Results {
		id := int(searchResult.ID)

		name := searchResult.Title
		if searchResult.Title != searchResult.OriginalTitle {
			name = name + " (" + searchResult.OriginalTitle + ")"
		}

		year, _ := strconv.Atoi(searchResult.ReleaseDate[:4])

		credits, _ := tmdbClient.GetMovieCredits(id, nil)
		directors := []string{}
		for _, person := range credits.Crew {
			if person.Job == "Director" {
				directors = append(directors, person.Name)
			}
		}

		var currMovie = &basicMovie{
			ID:        strconv.Itoa(id),
			Name:      name,
			Year:      year,
			Directors: directors,
			Poster:    tmdb.GetImageURL(searchResult.PosterPath, tmdb.Original),
		}

		results = append(results, *currMovie)
	}

	c.IndentedJSON(http.StatusOK, results)
}

func log(c *gin.Context) {
	var m db.Movie
	if jsonParseErr := c.BindJSON(&m); jsonParseErr != nil {
		panic(jsonParseErr)
	}

	db.InitNotion()
	db.Request(m.ToPageRequest())
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/search", search)
	router.POST("/log", log)

	router.Run("0.0.0.0:8080")
}
