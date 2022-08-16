package db

import (
	"encoding/json"
	"time"

	"github.com/jomei/notionapi"
)

type PageConverter interface {
	ConvertToProperty() notionapi.Property
}

type Title string

func (t Title) ConvertToProperty() notionapi.Property {
	return notionapi.TitleProperty{
		Title: []notionapi.RichText{
			{
				Text: notionapi.Text{
					Content: string(t),
				},
			},
		},
	}
}

type RichText string

func (t RichText) ConvertToProperty() notionapi.Property {
	return notionapi.RichTextProperty{
		RichText: []notionapi.RichText{
			{
				Text: notionapi.Text{
					Content: string(t),
				},
			},
		},
	}
}

type Number int

func (n Number) ConvertToProperty() notionapi.Property {
	return notionapi.NumberProperty{
		Number: float64(n),
	}
}

type MultiSelect []string

func (m MultiSelect) ConvertToProperty() notionapi.Property {
	options := []notionapi.Option{}
	for _, option := range m {
		options = append(options, notionapi.Option{
			Name: option,
		})
	}
	return notionapi.MultiSelectProperty{
		MultiSelect: options,
	}
}

type File string

func (f File) ConvertToProperty() notionapi.FilesProperty {
	return notionapi.FilesProperty{
		Files: []notionapi.File{
			{
				Type: "external",
				Name: string(f),
				External: &notionapi.FileObject{
					URL: string(f),
				},
			},
		},
	}
}

type Date time.Time

func (d Date) ConvertToProperty() notionapi.DateProperty {
	date := notionapi.Date(time.Time(d))
	return notionapi.DateProperty{
		Date: &notionapi.DateObject{
			Start: &date,
		},
	}
}

func (d *Date) UnmarshalJSON(p []byte) error {
	var s string
	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}
