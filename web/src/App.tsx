import { useState } from 'react'
import RatingForm from './components/RatingForm'
import SearchPage from './components/SearchPage'

export type Result = {
  id: ''
  name: string
  year: number
  directors: string[]
  poster: string
}

const DEFAULT_RESULT: Result = {
  id: '',
  name: '',
  year: 2022,
  directors: [],
  poster: '',
}

function App() {
  const [result, setResults] = useState(DEFAULT_RESULT)
  const [isResultSelected, setIsResultSelected] = useState(false)

  function handleSelection(result: Result) {
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
