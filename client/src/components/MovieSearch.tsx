import { SetStateAction, useEffect, useState } from "react"

export enum Form {
  Movie,
  Series,
}

type SearchProps = {
  onAutofill: (name: string, year: number, form: Form) => void
}

function MovieSearch({ onAutofill }: SearchProps) {
  const [nameInput, setNameInput] = useState("")
  const [yearInput, setYearInput] = useState(2022)
  const [formInput, setFormInput] = useState(Form.Movie)

  useEffect(() => {
    let timer = setTimeout(
      () => onAutofill(nameInput, yearInput, formInput),
      1000
    )
    return () => clearTimeout(timer)
  }, [yearInput, nameInput])

  useEffect(() => onAutofill(nameInput, yearInput, formInput), [formInput])

  const handleNameChange = (event: any) => {
    setNameInput(event.target.value)
  }

  const handleYearChange = (event: any) => {
    setYearInput(event.target.value)
  }

  const handleFormChange = (event: any) => {
    setFormInput(event.target.value)
  }

  return (
    <>
      <input type="text" value={nameInput} onChange={handleNameChange} />
      <input type="number" value={yearInput} onChange={handleYearChange} />
      <select value={formInput} onChange={handleFormChange}>
        <option value={Form.Movie}>Movie</option>
        <option value={Form.Series}>Series</option>
      </select>
    </>
  )
}

export default MovieSearch
