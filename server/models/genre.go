package models

type Genre string

const (
	genreAction     Genre = "Action"
	genreAdventure  Genre = "Adventure"
	genreComedy     Genre = "Comedy"
	genreDarkComedy Genre = "Dark Comedy"
	genreDrama      Genre = "Drama"
	genreHorror     Genre = "Horror"
	genreSlasher    Genre = "Slasher"
	genreRomance    Genre = "Romance"
	genreThriller   Genre = "Thriller"
)

func genreOf(str string) Genre {
	switch str {
	case string(genreAction):
		return genreAction
	case string(genreAdventure):
		return genreAdventure
	case string(genreComedy):
		return genreComedy
	case string(genreDarkComedy):
		return genreDarkComedy
	case string(genreDrama):
		return genreDrama
	case string(genreHorror):
		return genreHorror
	case string(genreSlasher):
		return genreSlasher
	case string(genreRomance):
		return genreRomance
	case string(genreThriller):
		return genreThriller
	default:
		return ""
	}
}

type Genres []Genre

func (g Genres) ToStrings() []string {
	strings := []string{}
	for _, genre := range g {
		strings = append(strings, string(genre))
	}
	return strings
}

func (g Genres) IsEmpty() bool {
	return len(g.ToStrings()) == 0
}

type Theme string

const (
	themeComingOfAge Theme = "Coming of Age"
	themeCrime       Theme = "Crime"
	themeMystery     Theme = "Mystery"
	themeWar         Theme = "War"
)

func themeOf(str string) Theme {
	switch str {
	case string(themeComingOfAge):
		return themeComingOfAge
	case string(themeCrime):
		return themeCrime
	case string(themeMystery):
		return themeMystery
	case string(themeWar):
		return themeWar
	default:
		return ""
	}
}

type Themes []Theme

func (t Themes) ToStrings() []string {
	strings := []string{}
	for _, theme := range t {
		strings = append(strings, string(theme))
	}
	return strings
}

type Setting string

const (
	settingHistory   Setting = "History"
	settingFantasy   Setting = "Fantasy"
	settingSciFi     Setting = "SciFi"
	settingSuperhero Setting = "Superhero"
	settingWestern   Setting = "Western"
)

func settingOf(str string) Setting {
	if str == "Science Fiction" {
		str = string(settingSciFi)
	}
	switch str {
	case string(settingHistory):
		return settingHistory
	case string(settingFantasy):
		return settingFantasy
	case string(settingSciFi):
		return settingSciFi
	case string(settingSuperhero):
		return settingSuperhero
	case string(settingWestern):
		return settingWestern
	default:
		return ""
	}
}

type Settings []Setting

func (s Settings) ToStrings() []string {
	strings := []string{}
	for _, setting := range s {
		strings = append(strings, string(setting))
	}
	return strings
}

type Tag string

const (
	tagAnimated  Tag = "Animated"
	tagAnthology Tag = "Anthology"
	tagBiopic    Tag = "Biopic"
	tagFamily    Tag = "Family"
	tagMusical   Tag = "Musical"
)

func tagOf(str string) Tag {
	if str == "Biography" {
		str = "Biopic"
	}
	if str == "Music" {
		str = "Musical"
	}
	switch str {
	case string(tagAnimated):
		return tagAnimated
	case string(tagAnthology):
		return tagAnthology
	case string(tagBiopic):
		return tagBiopic
	case string(tagFamily):
		return tagFamily
	case string(tagMusical):
		return tagMusical
	default:
		return ""
	}
}

type Tags []Tag

func (t Tags) ToStrings() []string {
	strings := []string{}
	for _, tag := range t {
		strings = append(strings, string(tag))
	}
	return strings
}

func ConvertGenreStrings(strings []string) (genres Genres, themes Themes, settings Settings, tags Tags) {
	for _, str := range strings {
		if genre := genreOf(str); genre != "" {
			genres = append(genres, genre)
		}
		if theme := themeOf(str); theme != "" {
			themes = append(themes, theme)
		}
		if setting := settingOf(str); setting != "" {
			settings = append(settings, setting)
		}
		if tag := tagOf(str); tag != "" {
			tags = append(tags, tag)
		}
	}
	return
}
