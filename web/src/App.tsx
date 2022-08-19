import RatingForm from './components/RatingForm'
import SearchPage from './components/SearchPage'

function App() {
  const MOVIES = [
    {
      id: '760104',
      year: 2022,
      poster:
        'https://image.tmdb.org/t/p/original/woTQx9Q4b8aO13jR9dsj8C9JESy.jpg',
      title: 'X',
      directors: ['Ti West'],
    },
    // {
    //   poster:
    //     "https://image.tmdb.org/t/p/original/7IiTTgloJzvGI1TAYymCfbfl3vT.jpg",
    //   title: "Parasite (기생충)",
    //   director: "Bong Joon-ho",
    // },
  ]

  return (
    <>
      {/* <SearchPage /> */}
      <RatingForm
        id={MOVIES[0].id}
        name={MOVIES[0].title}
        year={MOVIES[0].year}
        directors={MOVIES[0].directors}
        poster={MOVIES[0].poster}
      />
    </>
  )
}

export default App
