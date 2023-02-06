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
    const watch = {
      movieID: result.ID,
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
      date: new Date(data.watchedDate).toISOString(),
    }
    fetch('http://localhost:8080/log', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(watch),
    })
      .then((response) => response.json)
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
