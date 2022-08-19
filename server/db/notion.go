package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jomei/notionapi"
)

var notion *notionapi.Client

func InitNotion() {
	notion = notionapi.NewClient(notionapi.Token(os.Getenv("NOTION_API_KEY")))
}

func Request(r *notionapi.PageCreateRequest) {
	p, err := notion.Page.Create(context.TODO(), r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(p.URL)
	}
}

type PageRequester interface {
	ToPageRequest() *notionapi.PageCreateRequest
}

func (m *Movie) ToPageRequest() *notionapi.PageCreateRequest {
	p := &notionapi.PageCreateRequest{}

	p.Parent = notionapi.Parent{
		DatabaseID: notionapi.DatabaseID(os.Getenv("NOTION_DB_ID")),
	}

	p.Properties = notionapi.Properties{
		"ID":         m.Id.ConvertToProperty(),
		"Title":      m.Title.ConvertToProperty(),
		"Year":       m.Year.ConvertToProperty(),
		"Director":   m.Directors.ConvertToProperty(),
		"Genre":      m.Genre.ConvertToProperty(),
		"Enjoyment":  m.Rating.Enjoyment.ConvertToProperty(),
		"Feel":       m.Rating.Feel.ConvertToProperty(),
		"Style":      m.Rating.Style.ConvertToProperty(),
		"Characters": m.Rating.Characters.ConvertToProperty(),
		"Narrative":  m.Rating.Narrative.ConvertToProperty(),
		"Impact":     m.Rating.Impact.ConvertToProperty(),
		"Interest":   m.Rating.Interest.ConvertToProperty(),
		"Poster":     m.Poster.ConvertToProperty(),
		"Watched":    m.Watched.ConvertToProperty(),
	}

	return p
}
