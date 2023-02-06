package notion

import (
	"time"

	no "github.com/dstotijn/go-notion"
)

func createTitleProperty(title string) no.DatabasePageProperty {
	return no.DatabasePageProperty{
		Title: []no.RichText{
			{
				Text: &no.Text{
					Content: title,
				},
			},
		},
	}
}

func createNumberProperty(number uint) no.DatabasePageProperty {
	convertedNumber := float64(number)
	return no.DatabasePageProperty{
		Number: &convertedNumber,
	}
}

func addMultiselectPropery(name string, selected []string, props no.DatabasePageProperties) no.DatabasePageProperties {
	if len(selected) == 0 {
		return props
	}
	options := []no.SelectOptions{}
	for _, option := range selected {
		options = append(options, no.SelectOptions{
			Name: option,
		})
	}
	props[name] = no.DatabasePageProperty{
		MultiSelect: options,
	}
	return props
}

func createFileProperty(url string) no.DatabasePageProperty {
	return no.DatabasePageProperty{
		Files: []no.File{
			{
				Name: "name",
				Type: no.FileTypeExternal,
				External: &no.FileExternal{
					URL: url,
				},
			},
		},
	}
}

func createDateProperty(date time.Time) no.DatabasePageProperty {
	return no.DatabasePageProperty{
		Date: &no.Date{
			Start: no.NewDateTime(date, false),
		},
	}
}

func createRelationProperty(pageID string) no.DatabasePageProperty {
	return no.DatabasePageProperty{
		Relation: []no.Relation{
			{
				ID: pageID,
			},
		},
	}
}
