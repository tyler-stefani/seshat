interface MovieProps {
  poster: string
  title: String
  director: String
}

function MovieInfo(props: MovieProps) {
  return (
    <div className="w-64">
      <img src={props.poster} className="w-full"></img>
      <h2>{props.title}</h2>
      <p>{props.director}</p>
    </div>
  )
}

export default MovieInfo
