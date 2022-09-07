import { Result } from '../App'

interface MovieProps {
  result: Result
  onClick: (result: Result) => void
}

function MovieInfo({ result, onClick }: MovieProps) {
  function handleSelection() {
    onClick(result)
  }

  return (
    <div className='w-64' onClick={handleSelection}>
      <img src={result.poster} className='w-full'></img>
      <h2>{result.name}</h2>
      <p>{result.directors.length > 0 ? result.directors[0] : ''}</p>
    </div>
  )
}

export default MovieInfo
