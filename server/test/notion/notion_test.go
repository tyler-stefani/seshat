package notion_test

import (
	"context"
	"seshat-server/models"
	"seshat-server/notion"
	"testing"

	no "github.com/dstotijn/go-notion"
	"github.com/golang/mock/gomock"
)

func TestWriteMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientMock := NewMockClient(ctrl)

	movieDatabaseID := "dbid"

	id := 0
	title := "Funnee Movie"
	year := 2023
	director := []string{"Peter Chumpus", "Paulie Chumpus"}
	posterurl := "funneemovie.jpg"
	genres := []models.Genre{"Comedy"}

	params := no.CreatePageParams{
		ParentType: no.ParentTypeDatabase,
		ParentID:   movieDatabaseID,
		DatabasePageProperties: &no.DatabasePageProperties{
			"ID": no.DatabasePageProperty{
				Title: []no.RichText{
					{
						Text: &no.Text{
							Content: title,
						},
					},
				},
			},
		},
	}
	clientMock.EXPECT().CreatePage(context.Background(), params).Return(no.Page{}, nil)

	movie := models.Movie{
		ID:        uint(id),
		Title:     title,
		Year:      uint(year),
		Director:  director,
		PosterURL: posterurl,
		Genre:     models.Genres(genres),
	}

	notion := notion.Notion{}
	notion.Client = clientMock
	notion.WriteMovie(movie, movieDatabaseID)
}
