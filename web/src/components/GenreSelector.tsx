import { UseFormRegister } from 'react-hook-form'
import { FormValues } from './RatingForm'

const GENRES = [
  'Action',
  'Adventure',
  'Animated',
  'Anthology',
  'Biopic',
  'Comedy',
  'Coming of Age',
  'Crime',
  'Dark Comedy',
  'Documentary',
  'Drama',
  'Family',
  'Fantasy',
  'History',
  'Horror',
  'Musical',
  'Mystery',
  'Romance',
  'SciFi',
  'Slasher',
  'Superhero',
  'Thriller',
  'War',
  'Western',
]

type GenreSelectorProps = {
  register: UseFormRegister<FormValues>
}

export default function GenreSelector({ register }: GenreSelectorProps) {
  return (
    <label className='block text-left w-64'>
      <span className='text-gray-700'>Genre</span>
      <select
        {...register('genre')}
        className='border-slate-400 border-2 rounded form-multiselect block w-full my-1 '
        id='genre'
        name='genre'
        multiple
      >
        {GENRES.map((genre) => {
          return <option value={genre}>{genre}</option>
        })}
      </select>
    </label>
  )
}
