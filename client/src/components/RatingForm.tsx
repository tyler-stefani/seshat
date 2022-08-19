import { useForm } from 'react-hook-form'
import GenreSelector from './GenreSelector'
import MovieInfo from './MovieInfo'
import RatingCategory from './RatingCategory'
import WatchedDateInput from './WatchedDateInput'

type RatingFormProps = {
  id: string
  name: string
  year: number
  directors: string[]
  poster: string
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

export default function RatingForm({
  id,
  name,
  year,
  directors,
  poster,
}: RatingFormProps) {
  const { register, handleSubmit } = useForm<FormValues>()

  const onSubmit = (data: FormValues) => {
    const movie = {
      id: id,
      title: name,
      year: year,
      poster: poster,
      directors: directors,
      rating: {
        enjoyment: parseInt(data.enjoyment),
        feel: parseInt(data.feel),
        style: parseInt(data.style),
        chatacters: parseInt(data.characters),
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
  }

  return (
    <form className="flex flex-col p-2" onSubmit={handleSubmit(onSubmit)}>
      <MovieInfo poster={poster} title={name} director={directors[0]} />
      <WatchedDateInput register={register} />
      {CATEGORIES.map((category) => (
        <RatingCategory register={register} category={category} />
      ))}
      <GenreSelector register={register} />
      <input type="submit" value="Submit" />
    </form>
  )
}
