package services

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/ports/repositories"
	"time"
)

type MovieService struct {
	movieRepo repositories.MovieRepository
}

func NewMovieService(movieRepo repositories.MovieRepository) *MovieService {
	return &MovieService{movieRepo}
}

func (ms *MovieService) CreateMovie(title, description string, releaseDate time.Time, rating uint8) (*domain.Movie, error) {
	return ms.movieRepo.CreateMovie(title, description, releaseDate, rating)
}

func (ms *MovieService) GetMovie(id int) (*domain.Movie, error) {
	return ms.movieRepo.GetMovie(id)
}

func (ms *MovieService) GetMovies(sortKey, sortOrder, title, actor string) ([]*domain.Movie, error) {
	if len(sortKey) == 0 {
		sortKey = "rating"
	}
	if len(sortOrder) == 0 {
		sortOrder = "desc"
	}
	if len(actor) == 0 {
		if len(title) == 0 {
			return ms.movieRepo.GetMovies()
		}
		return ms.movieRepo.SearchMoviesByTitle(sortKey, sortOrder, title)
	}

	return ms.movieRepo.SearchMovies(sortKey, sortOrder, title, actor)
}

func (ms *MovieService) UpdateMovie(id int, title, description string, releaseDate time.Time, rating uint8) (*domain.Movie, error) {
	return ms.movieRepo.UpdateMovie(id, title, description, releaseDate, rating)
}

func (ms *MovieService) DeleteMovie(id int) (int, error) {
	err := ms.movieRepo.DeleteMovie(id)

	return id, err
}
