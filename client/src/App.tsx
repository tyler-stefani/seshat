import MovieInfo from "./components/MovieInfo"
import MovieSearch from "./components/MovieSearch"
import RatingForm from "./components/RatingForm"
import SearchPage from "./components/SearchPage"

function App() {
  const MOVIES = [
    {
      poster:
        "https://image.tmdb.org/t/p/original/woTQx9Q4b8aO13jR9dsj8C9JESy.jpg",
      title: "X",
      director: "Ti West",
    },
    {
      poster:
        "https://image.tmdb.org/t/p/original/7IiTTgloJzvGI1TAYymCfbfl3vT.jpg",
      title: "Parasite (기생충)",
      director: "Bong Joon-ho",
    },
  ]

  return (
    <div>
      {/* {MOVIES.map((movie) => (
        <MovieInfo
          poster={movie.poster}
          title={movie.title}
          director={movie.director}
        ></MovieInfo>
      ))} */}
      <SearchPage />
    </div>
  )
}

export default App
