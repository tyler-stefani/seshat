import { UseFormRegister } from "react-hook-form"
import { FormValues } from "./RatingForm"

type DateInputProps = {
  register: UseFormRegister<FormValues>
}

export default function WatchedDateInput({ register }: DateInputProps) {
  return (
    <input
      {...register("watchedDate")}
      id="watchedDate"
      name="watchedDate"
      type="date"
    />
  )
}
