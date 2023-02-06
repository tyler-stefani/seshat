import { useState } from 'react'
import RatingForm from './components/RatingForm'
import SearchPage from './components/SearchPage'

export type Result = {
  ID: number
  Title: string
  Year: number
  Director: string[]
  PosterURL: string
  Genre: string[]
}

const DEFAULT_RESULT: Result = {
  ID: 0,
  Title: '',
  Year: 2022,
  Director: [],
  PosterURL: '',
  Genre: [],
}

function App() {
  const [result, setResults] = useState(DEFAULT_RESULT)
  const [isResultSelected, setIsResultSelected] = useState(false)

  function handleSelection(result: Result) {
    fetch('http://localhost:8080/add-movie', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(result),
    })
      .then((response) => response.json)
      .catch((error) => {
        console.error('Error:', error)
      })
    setResults(result)
    setIsResultSelected(true)
  }

  function handleLog() {
    setResults(DEFAULT_RESULT)
    setIsResultSelected(false)
  }

  return (
    <>
      {isResultSelected ? (
        <RatingForm result={result} onLog={handleLog} />
      ) : (
        <SearchPage onSelect={handleSelection} />
      )}
    </>
  )
}

export default App
