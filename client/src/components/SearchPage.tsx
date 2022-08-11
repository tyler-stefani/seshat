import { useState } from "react"
import MovieInfo from "./MovieInfo"
import MovieSearch, { Form } from "./MovieSearch"

type Result = {
  name: string
  year: number
  directors: string[]
  poster: string
}

export default function SearchPage() {
  const [searchResults, setSearchResults] = useState<Result[]>([])

  const getOptions = (name: string, year: number, form: Form) => {
    if (name === "" && year === 2022 && form === Form.Movie) {
    } else {
      const url = new URL("http://localhost:8080/search")
      url.searchParams.append("name", name)
      url.searchParams.append("year", "" + year)
      url.searchParams.append("form", form.toString())

      fetch(url, {
        headers: {
          "Content-Type": "application/json",
        },
      })
        .then((response) => response.json())
        .then((movies) => {
          setSearchResults(movies)
        })
    }
  }

  return (
    <>
      <MovieSearch onAutofill={getOptions} />
      {searchResults.map((result) => (
        <MovieInfo
          poster={result.poster}
          title={result.name}
          director={result.directors.length > 0 ? result.directors[0] : ""}
        />
      ))}
    </>
  )
}
