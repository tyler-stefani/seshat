package main

import (
	http "net/http"
	"os"

	"seshat-server/models"
	"seshat-server/notion"
	"seshat-server/tmdb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func search(c *gin.Context) {
	searchName := c.Query("name")
	searchYear := c.Query("year")

	search := tmdb.TMDB{}
	search.Connect()
	results := search.Search(searchName, searchYear)
	c.IndentedJSON(http.StatusOK, results)
}

func addMovie(c *gin.Context) {
	var mp models.MovieProps
	if jsonParseErr := c.BindJSON(&mp); jsonParseErr != nil {
		panic(jsonParseErr)
	}

	db := notion.Notion{}
	db.Connect()

	m := models.PropsToMovie(mp)
	db.WriteMovie(m, os.Getenv("MOVIE_DB_ID"))
}

func log(c *gin.Context) {
	var wp models.WatchProps
	if jsonParseErr := c.BindJSON(&wp); jsonParseErr != nil {
		panic(jsonParseErr)
	}

	db := notion.Notion{}
	db.Connect()

	w := models.PropsToWatch(wp)
	db.LogWatch(w, os.Getenv("WATCH_DB_ID"), os.Getenv("MOVIE_DB_ID"))
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/search", search)
	router.POST("/add-movie", addMovie)
	router.POST("/log", log)

	router.Run("0.0.0.0:8080")
}
