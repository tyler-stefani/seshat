import { useForm } from 'react-hook-form'
import { Result } from '../App'
import GenreSelector from './GenreSelector'
import MovieInfo from './MovieInfo'
import RatingCategory from './RatingCategory'
import WatchedDateInput from './WatchedDateInput'

type RatingFormProps = {
  result: Result
  onLog: () => void
}

export type FormValues = {
  watchedDate: Date
  enjoyment: string
  feel: string
  style: string
  characters: string
  narrative: string
  impact: string
  interest: string
  genre: string[]
}

const CATEGORIES = [
  'Enjoyment',
  'Feel',
  'Style',
  'Characters',
  'Narrative',
  'Impact',
  'Interest',
]

export default function RatingForm({ result, onLog }: RatingFormProps) {
  const { register, handleSubmit } = useForm<FormValues>()

  const onSubmit = (data: FormValues) => {
    const movie = {
      id: result.id,
      title: result.name,
      year: result.year,
      poster: result.poster,
      directors: result.directors,
      rating: {
        enjoyment: parseInt(data.enjoyment),
        feel: parseInt(data.feel),
        style: parseInt(data.style),
        characters: parseInt(data.characters),
        narrative: parseInt(data.narrative),
        impact: parseInt(data.impact),
        interest: parseInt(data.interest),
      },
      genre: data.genre,
      watched: data.watchedDate,
    }
    fetch('http://localhost:8080/log', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(movie),
    })
      .then((response) => response.json)
      .then((_) => {
        console.log('Success:')
      })
      .catch((error) => {
        console.error('Error:', error)
      })
    onLog()
  }

  return (
    <form className='flex flex-col p-2' onSubmit={handleSubmit(onSubmit)}>
      <MovieInfo result={result} onClick={() => {}} />
      <WatchedDateInput register={register} />
      {CATEGORIES.map((category) => (
        <RatingCategory register={register} category={category} />
      ))}
      <GenreSelector register={register} />
      <input type='submit' value='Submit' />
    </form>
  )
}
