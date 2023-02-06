package notion

import (
	"context"
	"os"
	"seshat-server/models"

	no "github.com/dstotijn/go-notion"
)

type Client interface {
	CreatePage(ctx context.Context, params no.CreatePageParams) (page no.Page, err error)
	UpdatePage(ctx context.Context, pageID string, params no.UpdatePageParams) (page no.Page, err error)
	QueryDatabase(ctx context.Context, id string, query *no.DatabaseQuery) (result no.DatabaseQueryResponse, err error)
}

type Notion struct {
	Client Client
}

func (n *Notion) Connect() {
	n.Client = no.NewClient(os.Getenv("NOTION_API_KEY"))
}

func (n *Notion) WriteMovie(m models.Movie, databaseID string) {
	if pageID := n.getMoviePageID(databaseID, m.ID); pageID == "" {
		params := movieToNotion(m)
		params.ParentID = databaseID
		_, err := n.Client.CreatePage(context.Background(), params)
		if err != nil {
			panic(err)
		}
	}
}

func movieToNotion(m models.Movie) no.CreatePageParams {
	databasePageProperties := no.DatabasePageProperties{
		"ID":     createNumberProperty(m.ID),
		"Title":  createTitleProperty(m.Title),
		"Year":   createNumberProperty(m.Year),
		"Poster": createFileProperty(m.PosterURL),
	}
	databasePageProperties = addMultiselectPropery("Director", m.Director, databasePageProperties)
	databasePageProperties = addMultiselectPropery("Genre", m.Genre.ToStrings(), databasePageProperties)
	databasePageProperties = addMultiselectPropery("Theme", m.Theme.ToStrings(), databasePageProperties)
	databasePageProperties = addMultiselectPropery("Setting", m.Setting.ToStrings(), databasePageProperties)
	databasePageProperties = addMultiselectPropery("Tag", m.Tag.ToStrings(), databasePageProperties)

	return no.CreatePageParams{
		ParentType: no.ParentTypeDatabase,
		Icon: &no.Icon{
			Type: no.IconTypeExternal,
			External: &no.FileExternal{
				URL: MOVIE_ICON_URL,
			},
		},
		DatabasePageProperties: &databasePageProperties,
	}
}

func (n *Notion) LogWatch(w models.Watch, watchDatabaseID string, movieDatabaseID string) {
	moviePageID := n.getMoviePageID(movieDatabaseID, w.MovieID)
	params := watchToNotion(w, moviePageID)
	params.ParentID = watchDatabaseID
	_, err := n.Client.CreatePage(context.Background(), params)
	if err != nil {
		panic(err)
	}
}

func watchToNotion(w models.Watch, pageID string) no.CreatePageParams {
	databasePageProperties := no.DatabasePageProperties{
		"Movie":      createRelationProperty(pageID),
		"Enjoyment":  createNumberProperty(uint(w.Rating.Enjoyment)),
		"Feel":       createNumberProperty(uint(w.Rating.Feel)),
		"Style":      createNumberProperty(uint(w.Rating.Style)),
		"Characters": createNumberProperty(uint(w.Rating.Characters)),
		"Narrative":  createNumberProperty(uint(w.Rating.Narrative)),
		"Impact":     createNumberProperty(uint(w.Rating.Impact)),
		"Interest":   createNumberProperty(uint(w.Rating.Interest)),
	}
	databasePageProperties = addMultiselectPropery("Genre", w.Genre.ToStrings(), databasePageProperties)
	databasePageProperties = addMultiselectPropery("Theme", w.Theme.ToStrings(), databasePageProperties)
	databasePageProperties = addMultiselectPropery("Setting", w.Setting.ToStrings(), databasePageProperties)
	databasePageProperties = addMultiselectPropery("Tag", w.Tag.ToStrings(), databasePageProperties)
	if !w.Date.IsZero() {
		(map[string]no.DatabasePageProperty)(databasePageProperties)["Date"] = createDateProperty(w.Date)
	}
	return no.CreatePageParams{
		ParentType: no.ParentTypeDatabase,
		Icon: &no.Icon{
			Type: no.IconTypeExternal,
			External: &no.FileExternal{
				URL: WATCH_ICON_URL,
			},
		},
		DatabasePageProperties: &databasePageProperties,
	}
}

func (n *Notion) getMoviePageID(databaseID string, id uint) (pageID string) {
	intID := int(id)
	result, err := n.Client.QueryDatabase(context.Background(), databaseID, &no.DatabaseQuery{
		Filter: &no.DatabaseQueryFilter{
			Property: "ID",
			DatabaseQueryPropertyFilter: no.DatabaseQueryPropertyFilter{
				Number: &no.NumberDatabaseQueryFilter{
					Equals: &intID,
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	if len(result.Results) == 0 {
		return
	} else {
		pageID = result.Results[0].ID
		return
	}
}

func (n *Notion) readAllMovies(databaseID string) map[string]models.Movie {
	results := make(map[string]models.Movie)
	query := no.DatabaseQuery{
		Sorts: []no.DatabaseQuerySort{
			{
				Property:  "Title",
				Direction: no.SortDirAsc,
			},
		},
	}
	result, err := n.Client.QueryDatabase(context.Background(), databaseID, &query)
	if err != nil {
		panic(err)
	}
	for _, page := range result.Results {
		movie := readMovieFromPage(page)
		results[page.ID] = movie
	}
	for result.HasMore {
		query.StartCursor = *result.NextCursor
		result, err = n.Client.QueryDatabase(context.Background(), databaseID, &query)
		if err != nil {
			panic(err)
		}
		for _, page := range result.Results {
			movie := readMovieFromPage(page)
			results[page.ID] = movie
		}
	}
	return results
}

func readMovieFromPage(p no.Page) (m models.Movie) {
	props := p.Properties.(no.DatabasePageProperties)
	m.ID = uint(*props["ID"].Number)
	m.Title = props["Title"].Title[0].Text.Content
	m.Year = uint(*props["Year"].Number)
	for _, option := range props["Director"].MultiSelect {
		m.Director = append(m.Director, option.Name)
	}
	for _, option := range props["Genre"].MultiSelect {
		m.Genre = append(m.Genre, models.Genre(option.Name))
	}
	m.PosterURL = props["Poster"].Files[0].External.URL
	return
}
