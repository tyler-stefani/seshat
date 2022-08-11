import { useForm } from "react-hook-form"
import GenreSelector from "./GenreSelector"
import RatingCategory from "./RatingCategory"
import WatchedDateInput from "./WatchedDateInput"

export type FormValues = {
  watchedDate: Date
  enjoyment: number
  feel: number
  style: number
  characters: number
  narrative: number
  impact: number
  interest: number
  genre: string[]
}

const CATEGORIES = [
  "Enjoyment",
  "Feel",
  "Style",
  "Characters",
  "Narrative",
  "Inpact",
  "Interest",
]

export default function RatingForm() {
  const { register, handleSubmit } = useForm<FormValues>()

  const onSubmit = (data: any) => {
    console.log(data)
  }

  return (
    <form className="px-8 pt-6 pb-8" onSubmit={handleSubmit(onSubmit)}>
      <WatchedDateInput register={register} />
      {CATEGORIES.map((category) => (
        <RatingCategory register={register} category={category} />
      ))}
      <GenreSelector register={register} />
      <input type="submit" value="Submit" />
    </form>
  )
}
