import { useState } from 'react'
import { Result } from '../App'
import MovieInfo from './MovieInfo'
import MovieSearch, { Form } from './MovieSearch'

type SearchPageProps = {
  onSelect: (result: Result) => void
}

export default function SearchPage({ onSelect }: SearchPageProps) {
  const [searchResults, setSearchResults] = useState<Result[]>([])

  const getOptions = (name: string, year: number, form: Form) => {
    if (name === '' && year === 2022 && form === Form.Movie) {
    } else {
      const url = new URL('http://localhost:8080/search')
      url.searchParams.append('name', name)
      url.searchParams.append('year', '' + year)
      url.searchParams.append('form', form.toString())

      fetch(url, {
        headers: {
          'Content-Type': 'application/json',
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
        <MovieInfo onClick={onSelect} result={result} />
      ))}
    </>
  )
}
