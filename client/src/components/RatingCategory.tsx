import { UseFormRegister } from "react-hook-form"
import { FormValues } from "./RatingForm"
import RatingScoreInput from "./RatingScoreInput"

type RatingCategoryProps = {
  register: UseFormRegister<FormValues>
  category: string
}

export default function RatingCategory({
  register,
  category,
}: RatingCategoryProps) {
  return (
    <div>
      <p>{category}</p>
      {[...range(0, 5)].map((score) => (
        <RatingScoreInput
          register={register}
          category={category}
          score={score}
        />
      ))}
    </div>
  )
}

function range(start: number, end: number) {
  var ans = []
  for (let i = start; i <= end; i++) {
    ans.push(i)
  }
  return ans
}
