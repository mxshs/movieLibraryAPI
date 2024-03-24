package services

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/ports/repositories"
	"sync"

	"time"
)

type MovieService struct {
	movieRepo      repositories.MovieRepository
	movieActorRepo repositories.MovieActorRepository
}

func NewMovieService(movieRepo repositories.MovieRepository, movieActorRepo repositories.MovieActorRepository) *MovieService {
	return &MovieService{
		movieRepo,
		movieActorRepo,
	}
}

func (ms *MovieService) CreateMovie(title, description string, releaseDate time.Time, rating uint8) (*domain.Movie, error) {
	return ms.movieRepo.CreateMovie(title, description, releaseDate, rating)
}

func (ms *MovieService) GetMovie(id int) (*domain.MovieDetail, error) {
	movie, err := ms.movieRepo.GetMovie(id)
	if err != nil {
		return nil, err
	}

	actors, err := ms.movieActorRepo.GetMovieActors(id)
	if err != nil {
		return nil, err
	}

	movie.Stars = actors

	return movie, nil
}

func (ms *MovieService) GetMovies(sortKey, sortOrder, title, actor string) ([]*domain.MovieDetail, error) {
	if len(sortKey) == 0 {
		sortKey = "rating"
	}
	if len(sortOrder) == 0 {
		sortOrder = "desc"
	}

	var movies []*domain.MovieDetail
	var err error

	switch len(title) + len(actor) {
	case 0:
		movies, err = ms.movieRepo.GetMovies()
		if err != nil {
			return nil, err
		}
	default:
		movies, err = ms.movieRepo.SearchMovies(sortKey, sortOrder, title, actor)
		if err != nil {
			return nil, err
		}
	}

	var wg sync.WaitGroup

	for i := range len(movies) {
		wg.Add(1)
		go func(id int) {
			actors, err := ms.movieActorRepo.GetMovieActors(id)
			if err != nil {
				return
			}

			movies[i].Stars = actors
			wg.Done()
		}(i)
	}

	wg.Wait()

	return movies, nil
}

func (ms *MovieService) UpdateMovie(id int, title, description string, releaseDate time.Time, rating uint8) (*domain.Movie, error) {
	return ms.movieRepo.UpdateMovie(id, title, description, releaseDate, rating)
}

func (ms *MovieService) DeleteMovie(id int) (int, error) {
	err := ms.movieRepo.DeleteMovie(id)

	return id, err
}
