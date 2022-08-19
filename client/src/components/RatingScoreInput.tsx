import { UseFormRegister } from 'react-hook-form'
import { FormValues } from './RatingForm'

type RatingScoreInputProps = {
  register: UseFormRegister<FormValues>
  category: string
  score: number
}

type Category =
  | 'enjoyment'
  | 'feel'
  | 'style'
  | 'characters'
  | 'narrative'
  | 'impact'
  | 'interest'

export default function RatingScoreInput({
  register,
  category,
  score,
}: RatingScoreInputProps) {
  return (
    <label htmlFor={category.toLowerCase + '-' + score.toString}>
      <input
        {...register(category.toLowerCase() as Category, {
          valueAsNumber: true,
        })}
        type='radio'
        name={category.toLowerCase()}
        value={score}
        id={category.toLowerCase + '-' + score.toString}
      />
      {score}
    </label>
  )
}
