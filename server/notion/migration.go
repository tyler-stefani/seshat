package notion

const OVERALL_FORMULA = "round((prop(\"Enjoyment\") * 5 + prop(\"Feel\") * 4 + prop(\"Style\") * 3 + prop(\"Characters\") * 3 + prop(\"Impact\") * 3 + prop(\"Narrative\") * 2 + prop(\"Interest\") * 2) / 22 * 100) / 100"
const STARS_FORMULA = "(prop(\"Overall\") < 0.5) ? 0.5 : (round(prop(\"Overall\") * 2) / 2)"
const MOON_FORMULA = "concat((prop(\"Stars\") < 0.5) ? \"🌑\" : ((prop(\"Stars\") == 0.5) ? \"🌗\" : \"🌕\"), (prop(\"Stars\") < 1.5) ? \"🌑\" : ((prop(\"Stars\") == 1.5) ? \"🌗\" : \"🌕\"), (prop(\"Stars\") < 2.5) ? \"🌑\" : ((prop(\"Stars\") == 2.5) ? \"🌗\" : \"🌕\"), (prop(\"Stars\") < 3.5) ? \"🌑\" : ((prop(\"Stars\") == 3.5) ? \"🌗\" : \"🌕\"), (prop(\"Stars\") < 4.5) ? \"🌑\" : ((prop(\"Stars\") == 4.5) ? \"🌗\" : \"🌕\"))"

const CAMERA_ICON_URL = "https://www.notion.so/icons/movie-camera_lightgray.svg"
const MOVIE_ICON_URL = "https://www.notion.so/icons/movie_gray.svg"
const WATCH_ICON_URL = "https://www.notion.so/icons/playback-play_gray.svg"
const TV_ICON_URL = "https://www.notion.so/icons/tv_gray.svg"

// func (n *Notion) GetDatabase(id string) {
// 	db, err := n.client.FindDatabaseByID(context.Background(), id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(db.Properties["Director"].MultiSelect.Options)
// }

// func (n *Notion) CreateDatabase(parentID string) string {
// 	db, err := n.client.CreateDatabase(context.Background(), no.CreateDatabaseParams{
// 		ParentPageID: parentID,
// 		Title: []no.RichText{
// 			{
// 				Text: &no.Text{
// 					Content: "Movies",
// 				},
// 			},
// 		},
// 		Icon: &no.Icon{
// 			Type: no.IconTypeExternal,
// 			External: &no.FileExternal{
// 				URL: CAMERA_ICON_URL,
// 			},
// 		},
// 		Properties: no.DatabaseProperties{
// 			"Title": no.DatabaseProperty{
// 				Type:  no.DBPropTypeTitle,
// 				Name:  "Title",
// 				Title: &no.EmptyMetadata{},
// 			},
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return db.ID
// }

// func (n *Notion) MigrateToVersionOne(databaseID string) {
// 	n.migrateBasicProps(databaseID)
// 	n.migrateUserProps(databaseID)
// 	n.migrateFormulaProps(databaseID)
// }

// func (n *Notion) migrateBasicProps(databaseID string) {
// 	basicProps := make(map[string]*no.DatabaseProperty)
// 	basicProps["ID"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "ID",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	basicProps["Title"] = &no.DatabaseProperty{
// 		Type:  no.DBPropTypeTitle,
// 		Name:  "Title",
// 		Title: &no.EmptyMetadata{},
// 	}
// 	basicProps["Year"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "Year",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	basicProps["Director"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeMultiSelect,
// 		Name: "Director",
// 		MultiSelect: &no.SelectMetadata{
// 			Options: []no.SelectOptions{},
// 		},
// 	}
// 	basicProps["Poster"] = &no.DatabaseProperty{
// 		Type:  no.DBPropTypeFiles,
// 		Name:  "Poster",
// 		Files: &no.EmptyMetadata{},
// 	}
// 	n.client.UpdateDatabase(context.Background(), databaseID, no.UpdateDatabaseParams{
// 		Properties: basicProps,
// 	})
// }

// func (n *Notion) migrateUserProps(databaseID string) {
// 	userProps := make(map[string]*no.DatabaseProperty)
// 	userProps["Watched"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeDate,
// 		Name: "Watched",
// 		Date: &no.EmptyMetadata{},
// 	}
// 	userProps["Enjoyment"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "Enjoyment",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	userProps["Feel"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "Feel",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	userProps["Style"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "Style",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	userProps["Characters"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "Characters",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	userProps["Narrative"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "Narrative",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	userProps["Impact"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "Impact",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	userProps["Interest"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeNumber,
// 		Name: "Interest",
// 		Number: &no.NumberMetadata{
// 			Format: no.NumberFormatNumber,
// 		},
// 	}
// 	userProps["Genre"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeMultiSelect,
// 		Name: "Genre",
// 		MultiSelect: &no.SelectMetadata{
// 			Options: []no.SelectOptions{
// 				{
// 					Name:  "Action",
// 					Color: no.ColorOrange,
// 				},
// 				{
// 					Name:  "Biopic",
// 					Color: no.ColorBrown,
// 				},
// 				{
// 					Name:  "Comedy",
// 					Color: no.ColorYellow,
// 				},
// 				{
// 					Name:  "Dark Comedy",
// 					Color: no.ColorYellow,
// 				},
// 				{
// 					Name:  "Documentary",
// 					Color: no.ColorBrown,
// 				},
// 				{
// 					Name:  "Drama",
// 					Color: no.ColorBlue,
// 				},
// 				{
// 					Name:  "Dramedy",
// 					Color: no.ColorBlue,
// 				},
// 				{
// 					Name:  "Family",
// 					Color: no.ColorPurple,
// 				},
// 				{
// 					Name:  "Fantasy",
// 					Color: no.ColorGreen,
// 				},
// 				{
// 					Name:  "Horror",
// 					Color: no.ColorRed,
// 				},
// 				{
// 					Name:  "Musical",
// 					Color: no.ColorPink,
// 				},
// 				{
// 					Name:  "Romance",
// 					Color: no.ColorPink,
// 				},
// 				{
// 					Name:  "RomCom",
// 					Color: no.ColorPink,
// 				},
// 				{
// 					Name:  "SciFi",
// 					Color: no.ColorGreen,
// 				},
// 				{
// 					Name:  "Superhero",
// 					Color: no.ColorPurple,
// 				},
// 				{
// 					Name:  "Thriller",
// 					Color: no.ColorRed,
// 				},
// 			},
// 		},
// 	}
// 	userProps["Overall"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeFormula,
// 		Name: "Overall",
// 		Formula: &no.FormulaMetadata{
// 			Expression: OVERALL_FORMULA,
// 		},
// 	}
// 	_, err := n.client.UpdateDatabase(context.Background(), databaseID, no.UpdateDatabaseParams{
// 		Properties: userProps,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (n *Notion) migrateFormulaProps(databaseID string) {
// 	formulaProps := make(map[string]*no.DatabaseProperty)
// 	formulaProps["Stars"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeFormula,
// 		Name: "Stars",
// 		Formula: &no.FormulaMetadata{
// 			Expression: STARS_FORMULA,
// 		},
// 	}
// 	_, err := n.client.UpdateDatabase(context.Background(), databaseID, no.UpdateDatabaseParams{
// 		Properties: formulaProps,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	formulaProps = make(map[string]*no.DatabaseProperty)
// 	formulaProps["☾"] = &no.DatabaseProperty{
// 		Type: no.DBPropTypeFormula,
// 		Name: "☾",
// 		Formula: &no.FormulaMetadata{
// 			Expression: MOON_FORMULA,
// 		},
// 	}
// 	_, err = n.client.UpdateDatabase(context.Background(), databaseID, no.UpdateDatabaseParams{
// 		Properties: formulaProps,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// }
