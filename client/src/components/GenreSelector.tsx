import { UseFormRegister } from "react-hook-form"
import { FormValues } from "./RatingForm"

const GENRES = [
  "Action",
  "BioPic",
  "Comedy",
  "Dark Comedy",
  "Documentary",
  "Drama",
  "Dramedy",
  "Family",
  "Fantasy",
  "Horror",
  "Musical",
  "Romance",
  "RomCom",
  "SciFi",
  "Superhero",
  "Thriller",
]

type GenreSelectorProps = {
  register: UseFormRegister<FormValues>
}

export default function GenreSelector({ register }: GenreSelectorProps) {
  return (
    <label className="block text-left w-64">
      <span className="text-gray-700">Genre</span>
      <select
        {...register("genre")}
        className="border-slate-400 border-2 rounded form-multiselect block w-full my-1 "
        id="genre"
        name="genre"
        multiple
      >
        {GENRES.map((genre) => {
          return <option value={genre.toLowerCase()}>{genre}</option>
        })}
      </select>
    </label>
  )
}
