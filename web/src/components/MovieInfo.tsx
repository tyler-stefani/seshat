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
      <img src={result.PosterURL} className='w-full'></img>
      <h2>{result.Title}</h2>
      <p>{result.Director.length > 0 ? result.Director[0] : ''}</p>
    </div>
  )
}

export default MovieInfo
